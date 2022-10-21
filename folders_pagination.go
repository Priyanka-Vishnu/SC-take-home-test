package folders

import "github.com/gofrs/uuid"

// Function name is GetAllFolders and it gets the requests else it shows error as the result
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	//Assigning variables to each
	//error is err
	//f1 is the Folder
	//fs is where it gets appended
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	// F is variable that is assigned to the folder and data is copied into it
	f := []Folder{}
	// r is a variable whch fetches all organization ID's requests
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	//for k and v r in range
	for k, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// Creating another function FetchAllFoldersByOrgID
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}

// Suggestion whould be to create 1 function under which various sub classes can be created.
// It can shorten the code as well as the code wouldrun faster when in sub class loops
