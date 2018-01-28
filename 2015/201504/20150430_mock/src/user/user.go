package user

// import "github.com/sgreben/testing-with-gomock/doer"
import "doer"

type User struct {
    Doer doer.Doer;
}

func (u *User) Use() error {
    return u.Doer.DoSomething(123, "Hello GoMock");
}
/*
//mockgen -destination=mocks/mock_doer.go -package=mocks github.com/sgreben/testing-with-gomock/doer Doer
mockgen -destination=mocks/mock_doer.go -package=mocks doer Doer
*/