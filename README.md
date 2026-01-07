# test

go set up
    go mod init <ชื่อ>
    go mod tidy


run test > go test -v ./...


go get -u

	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"


valid :
        requried, 
        type,
        email,
        url,
        matches(^[BMD]\\d{7}$),
		matches(^([0-3]\\.[0-9]{2}|4\\.00)$)~,
        stringlength(10|10),
		range(10|100)

