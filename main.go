package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	m "github.com/cjf93/bugzilla-data-importer/models"
	r "github.com/cjf93/bugzilla-data-importer/responses"
)

var basePath = "https://bugs.eclipse.org/bugs"

func main() {
	fmt.Println("Starting program")
	start := time.Now()
	fmt.Println("Starting time: " + start.String())
	fmt.Println("Get version")
	resp, err := GetVersion()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.Version)

	var profileMap m.ProfileMap
	profileMap = make(m.ProfileMap)

	var classificationMap m.ClassificationMap
	classificationMap = make(m.ClassificationMap)

	var componentMap m.ComponentMap
	componentMap = make(m.ComponentMap)

	var ProfileInserts []string
	var ProductInserts []string
	var ClassificationInserts []string
	var ComponentInserts []string
	var BugInserts []string

	//Get Product List
	ProductIDList, _ := GetProductsIDList()

	//For every Product
	for i := 0; i < len(ProductIDList.Ids); i++ {
		//Get Product for every ID from the Product List
		productID, _ := strconv.Atoi(ProductIDList.Ids[i])
		fmt.Println("Product #n: " + strconv.Itoa(i) + "/" + strconv.Itoa(len(ProductIDList.Ids)))
		//Query every product
		ProductQueryObject, _ := GetProductByID(productID)
		fmt.Println("ProductQueryObject: " + ProductIDList.Ids[i])

		//Create a Classificacion Object(int case it doesnt exists)
		var ClassificationObject m.Classification
		if _, ok := classificationMap[ProductQueryObject.Products[0].Classification]; !ok {

			//Get the Classificacion for every queryed product
			ClassificationQueryObject, _ := GetClassifiaciontByName(ProductQueryObject.Products[0].Classification)

			//Create a Classification Object
			ClassificationObject = m.CreateClassification(ClassificationQueryObject.Classifications[0].ID, ClassificationQueryObject.Classifications[0].Name)
			//Insert ClassificationObject into map
			classificationMap[ClassificationObject.Name] = ClassificationObject
			fmt.Println("ClassificationObject.ID: " + strconv.Itoa(ClassificationObject.ID))
			//Insert Classification Object
			insertS := ClassificationObject.GenerateInsert()

			//Put it on a []string
			ClassificationInserts = append(ClassificationInserts, insertS)

		}
		ClassificationObject = classificationMap[ProductQueryObject.Products[0].Classification]
		//Create a Product Object
		ProductObject := m.CreateProduct(productID, ClassificationObject.ID, ProductQueryObject.Products[0].Name, ProductQueryObject.Products[0].Description, ProductQueryObject.Products[0].DefaultMilestone, ProductQueryObject.Products[0].IsActive, ProductQueryObject.Products[0].HasUnconfirmed)
		fmt.Println("ClassificationObject.ID: " + strconv.Itoa(ClassificationObject.ID))
		//Insert a Product Object
		insertS := ProductObject.GenerateInsert()

		//TODO: Put it on a []string
		ProductInserts = append(ProductInserts, insertS)

		//For every Component of the Object
		for j := 0; j < len(ProductQueryObject.Products[0].Components); j++ {
			fmt.Println("Components: " + strconv.Itoa(ProductQueryObject.Products[0].Components[j].ID))
			//Create a Profile Object(int case it doesnt exists)
			var assignedToProfileObject m.Profile
			if _, ok := profileMap[ProductQueryObject.Products[0].Components[j].DefaultAssignedTo]; !ok {

				ProfileQuery, _ := GetProfileByName(ProductQueryObject.Products[0].Components[j].DefaultAssignedTo)
				assignedToProfileObject = m.CreateProfile(ProfileQuery.Users[0].ID, ProductQueryObject.Products[0].Components[j].DefaultAssignedTo, ProfileQuery.Users[0].RealName)

				//Insert Profile into map
				profileMap[assignedToProfileObject.Login_name] = assignedToProfileObject

				//Insert a Profile Object
				insertS := assignedToProfileObject.GenerateInsert()

				//Put it on a []string
				ProfileInserts = append(ProfileInserts, insertS)

			}
			assignedToProfileObject = profileMap[ProductQueryObject.Products[0].Components[j].DefaultAssignedTo]

			//Create a Profile Object(int case it doesnt exists)
			var qaContactProfileObject m.Profile
			var ComponentObject m.Component
			//First check if this field is empty(can be null in the DB)
			if ProductQueryObject.Products[0].Components[j].DefaultQaContact != "" {
				if qaContactProfileObject, ok := profileMap[ProductQueryObject.Products[0].Components[j].DefaultQaContact]; !ok {

					ProfileQuery, _ := GetProfileByName(ProductQueryObject.Products[0].Components[j].DefaultQaContact)
					qaContactProfileObject = m.CreateProfile(ProfileQuery.Users[0].ID, ProfileQuery.Users[0].Name, ProfileQuery.Users[0].RealName)

					//Insert Profile into map
					profileMap[qaContactProfileObject.Login_name] = qaContactProfileObject

					//Insert a Profile Object
					insertS := qaContactProfileObject.GenerateInsert()

					//Put it on a []string
					ProfileInserts = append(ProfileInserts, insertS)
				}
				qaContactProfileObject = profileMap[ProductQueryObject.Products[0].Components[j].DefaultQaContact]
				//Create a Component Object
				ComponentObject = m.CreateComponent(ProductQueryObject.Products[0].Components[j].ID, productID, assignedToProfileObject.Userid, qaContactProfileObject.Userid,
					ProductQueryObject.Products[0].Components[j].Name, ProductQueryObject.Products[0].Components[j].Description, ProductQueryObject.Products[0].Components[j].IsActive)

				//Insert a Component Object
				insertS := ComponentObject.GenerateInsert()
				componentMap[ComponentObject.Name] = ComponentObject

				//Put it on a []string
				ComponentInserts = append(ComponentInserts, insertS)

			} else {
				//Create a Component Object with null qa contact
				ComponentObject = m.CreateComponent(ProductQueryObject.Products[0].Components[j].ID, productID, assignedToProfileObject.Userid, -1,
					ProductQueryObject.Products[0].Components[j].Name, ProductQueryObject.Products[0].Components[j].Description, ProductQueryObject.Products[0].Components[j].IsActive)

				//Insert a Component Object
				insertS := ComponentObject.GenerateInsert()
				componentMap[ComponentObject.Name] = ComponentObject

				//Put it on a []string
				ComponentInserts = append(ComponentInserts, insertS)
			}
			if ProductQueryObject.Products[0].Components[j].ID == 1127 {
				fmt.Println(ProductQueryObject.Products[0].Components[j].ID)
				fmt.Println(ProductQueryObject.Products[0].Components[j].Name)
				fmt.Println(componentMap[ProductQueryObject.Products[0].Components[j].Name].ID)
			}
		}

		//Get all the bugs in the product
		BugQueryObject, _ := GetBugByProduct(ProductQueryObject.Products[0].Name)

		//For every Bug in the product
		for k := 0; k < len(BugQueryObject.Bugs); k++ {
			fmt.Println("SingleBug #n: " + strconv.Itoa(k) + "/" + strconv.Itoa(len(BugQueryObject.Bugs)))
			SingleBug := BugQueryObject.Bugs[k]
			fmt.Println("SingleBug: " + strconv.Itoa(SingleBug.ID))
			fmt.Println(SingleBug.Component)
			fmt.Println(componentMap[SingleBug.Component].ID)
			//Create a Profile Object(int case it doesnt exists) --- reporter
			var reporterProfileObject m.Profile
			if _, ok := profileMap[SingleBug.Creator]; !ok {
				reporterProfileObject = m.CreateProfile(SingleBug.CreatorDetail.ID, SingleBug.CreatorDetail.Email, SingleBug.CreatorDetail.RealName)

				//Insert Profile into map
				profileMap[reporterProfileObject.Login_name] = reporterProfileObject

				//Insert a Profile Object
				insertS := reporterProfileObject.GenerateInsert()

				//TODO: Put it on a []string
				ProfileInserts = append(ProfileInserts, insertS)
			}

			reporterProfileObject = profileMap[SingleBug.Creator]
			//Create a Profile Object(int case it doesnt exists) --- assigned
			var assignedToProfileObject m.Profile
			if _, ok := profileMap[SingleBug.AssignedTo]; !ok {

				assignedToProfileObject = m.CreateProfile(SingleBug.AssignedToDetail.ID, SingleBug.AssignedToDetail.Email, SingleBug.AssignedToDetail.RealName)

				//Insert Profile into map
				profileMap[assignedToProfileObject.Login_name] = assignedToProfileObject

				//Insert a Profile Object
				insertS := assignedToProfileObject.GenerateInsert()

				//TODO: Put it on a []string
				ProfileInserts = append(ProfileInserts, insertS)
			}

			assignedToProfileObject = profileMap[SingleBug.AssignedTo]
			//Create a Profile Object(int case it doesnt exists) --- qa_contact
			var qaContactProfileObject m.Profile
			//First check if this field is empty(can be null in the DB)
			if SingleBug.QaContact != "" {
				if qaContactProfileObject, ok := profileMap[SingleBug.QaContact]; !ok {

					qaProfileQuery, _ := GetProfileByName(SingleBug.QaContact)
					qaContactProfileObject = m.CreateProfile(qaProfileQuery.Users[0].ID, qaProfileQuery.Users[0].Name, qaProfileQuery.Users[0].RealName)

					//Insert Profile into map
					profileMap[qaContactProfileObject.Login_name] = qaContactProfileObject

					//Insert a Profile Object
					insertS := qaContactProfileObject.GenerateInsert()

					//TODO: Put it on a []string
					ProfileInserts = append(ProfileInserts, insertS)
				}
				qaContactProfileObject = profileMap[SingleBug.QaContact]

				//Create a Bug Object
				BugObject := m.CreateBug(SingleBug.ID, SingleBug.AssignedToDetail.ID, productID,
					SingleBug.CreatorDetail.ID, componentMap[SingleBug.Component].ID, qaContactProfileObject.Userid,
					SingleBug.Severity, SingleBug.Status, SingleBug.CreationTime.Format("2006-01-02 15:04:05"),
					SingleBug.Summary, SingleBug.OpSys, SingleBug.Priority, SingleBug.Platform,
					SingleBug.Version, SingleBug.Resolution, SingleBug.TargetMilestone, SingleBug.Whiteboard,
					SingleBug.LastChangeTime.Format("2006-01-02 15:04:05"), SingleBug.IsConfirmed, SingleBug.IsCreatorAccessible,
					SingleBug.IsCcAccessible)

				//Insert a Component Object
				insertS := BugObject.GenerateInsert()

				//TODO: Put it on a []string
				BugInserts = append(BugInserts, insertS)

			} else {
				//Create a Bug Object
				BugObject := m.CreateBug(SingleBug.ID, SingleBug.AssignedToDetail.ID, productID,
					SingleBug.CreatorDetail.ID, componentMap[SingleBug.Component].ID, -1,
					SingleBug.Severity, SingleBug.Status, SingleBug.CreationTime.Format("2006-01-02 15:04:05"),
					SingleBug.Summary, SingleBug.OpSys, SingleBug.Priority, SingleBug.Platform,
					SingleBug.Version, SingleBug.Resolution, SingleBug.TargetMilestone, SingleBug.Whiteboard,
					SingleBug.LastChangeTime.Format("2006-01-02 15:04:05"), SingleBug.IsConfirmed, SingleBug.IsCreatorAccessible,
					SingleBug.IsCcAccessible)

				//Insert a Component Object
				insertS := BugObject.GenerateInsert()

				//TODO: Put it on a []string
				BugInserts = append(BugInserts, insertS)
			}
		}
	}
	//Create an insert file for every object type (5)
	fmt.Println("PROFILE: " + strconv.Itoa(len(ProfileInserts)))
	WriteToFile("profileInserts", ProfileInserts)
	fmt.Println("CLASSIFICATION: " + strconv.Itoa(len(ClassificationInserts)))
	WriteToFile("classificationInserts", ClassificationInserts)
	fmt.Println("PRODUCT: " + strconv.Itoa(len(ProductInserts)))
	WriteToFile("productInserts", ProductInserts)
	fmt.Println("COMPONENT: " + strconv.Itoa(len(ComponentInserts)))
	WriteToFile("componentInserts", ComponentInserts)
	fmt.Println("BUG: " + strconv.Itoa(len(BugInserts)))
	WriteToFile("bugInserts", BugInserts)

	fmt.Println("Starting time: " + start.String())
	t := time.Now()
	fmt.Println("End time: " + t.String())
	elapsed := t.Sub(start)
	fmt.Println("Elapsed time: " + elapsed.String())
}

//GetVersion return a respones.GetVersion Type containing the HTTP response to the GET call /rest/version
func GetVersion() (r.GetVersion, error) {
	var response r.GetVersion

	resp, err := http.Get(basePath + "/rest/version")
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		var errorResponse r.ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println("Error on API response: " + errorResponse.Message)
		return response, errors.New("Error on API response: " + errorResponse.Message)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("Error on json.Unmarshall: " + err.Error())
	}

	return response, nil
}

//GetProductsIDList return a response.GetProductList Type containing the HTTP response to the GET call /rest/product_selectable
func GetProductsIDList() (r.GetProductList, error) {
	var response r.GetProductList

	resp, err := http.Get(basePath + "/rest/product_selectable")
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		var errorResponse r.ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println("Error on API response: " + errorResponse.Message)
		return response, errors.New("Error on API response: " + errorResponse.Message)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("Error on json.Unmarshall: " + err.Error())
	}

	return response, nil
}

//GetClassifiaciontByName return a response.GetBugs Type containing the HTTP response to the GET call /rest/classification/<name>
func GetClassifiaciontByName(name string) (r.GetClassification, error) {
	var response r.GetClassification

	resp, err := http.Get(basePath + "/rest/classification/" + name)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		var errorResponse r.ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println("Error on API response: " + errorResponse.Message)
		return response, errors.New("Error on API response: " + errorResponse.Message)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("Error on json.Unmarshall: " + err.Error())
	}

	return response, nil
}

//GetProductByID return a response.GetProducts Type containing the HTTP response to the GET call /rest/products?ids=id
func GetProductByID(id int) (r.GetProducts, error) { //The api can accept multiple ids. TODO: make it to query multiple ID's
	var response r.GetProducts

	resp, err := http.Get(basePath + "/rest/product?ids=" + strconv.Itoa(id))
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		var errorResponse r.ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println("Error on API response: " + errorResponse.Message)
		return response, errors.New("Error on API response: " + errorResponse.Message)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("Error on json.Unmarshall: " + err.Error())
	}

	return response, nil
}

//GetBugByID return a response.GetBugs Type containing the HTTP response to the GET call /rest/bug?id=id
func GetBugByID(id int) (r.GetBugs, error) {
	var response r.GetBugs

	resp, err := http.Get(basePath + "/rest/bug?id=" + strconv.Itoa(id))
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		var errorResponse r.ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println("Error on API response: " + errorResponse.Message)
		return response, errors.New("Error on API response: " + errorResponse.Message)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("Error on json.Unmarshall: " + err.Error())
	}

	return response, nil
}

//GetBugByProduct return a response.GetBugsv2 Type containing the HTTP response to the GET call /rest/bug?product=Foo
func GetBugByProduct(product string) (r.GetBugsv2, error) {
	var response r.GetBugsv2

	resp, err := http.Get(basePath + "/rest/bug?product=" + product)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		var errorResponse r.ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println("Error on API response: " + errorResponse.Message)
		return response, errors.New("Error on API response: " + errorResponse.Message)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("Error on json.Unmarshall: " + err.Error())
	}

	return response, nil
}

//GetProfileByName return a response.GetBugs Type containing the HTTP response to the GET call /user?names=<name>
func GetProfileByName(name string) (r.GetProfile, error) {
	var response r.GetProfile

	resp, err := http.Get(basePath + "/rest/user?names=" + url.QueryEscape(name))
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		var errorResponse r.ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println("Error on API response: " + errorResponse.Message)
		return response, errors.New("Error on API response: " + errorResponse.Message)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("Error on json.Unmarshall: " + err.Error())
	}

	return response, nil
}

func WriteToFile(fileName string, data []string) {
	file, err := os.OpenFile(fileName+".txt", os.O_RDWR, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range data {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
