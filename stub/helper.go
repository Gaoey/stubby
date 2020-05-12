package stub

import "github.com/hashicorp/go-uuid"

func GenerateID() string {
	id, _ := uuid.GenerateUUID()
	return id
}
