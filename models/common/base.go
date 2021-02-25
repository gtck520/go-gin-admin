package common

import(
	"fmt"

	"github.com/konger/ckgo/models/basemodel"
)

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", basemodel.GetTablePrefix(),"com_", name)
}