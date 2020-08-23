package global

import (
    "github.com/jinzhu/gorm"
    oplogging "github.com/op/go-logging"
)

var (
    GVA_DB     *gorm.DB
    GVA_LOG    *oplogging.Logger
)
