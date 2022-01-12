package factories

import (
    "huango/app/models/link"

    "github.com/bxcodec/faker/v3"
)

func MakeLinks(count int) []link.Link {

    var objs []link.Link

    for i := 0; i < count; i++ {
        linkModel := link.Link{
            Name: faker.Username(),
            URL:  faker.URL(),
        }
        objs = append(objs, linkModel)
    }

    return objs
}