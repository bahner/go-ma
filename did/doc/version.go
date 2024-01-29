package doc

import (
	"time"
)

func (d *Document) UpdateVersion() {

	d.Version = time.Now().Unix()

}
