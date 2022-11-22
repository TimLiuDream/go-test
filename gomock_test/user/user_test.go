package user

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/timliudream/go-test/gomock_test/mocks"
	"testing"
)

func TestUse(t *testing.T) {
	//core of gomock
	mockCtrl := gomock.NewController(t)
	//used to trigger final assertions. if its ignored, mocking assertions will never fail
	defer mockCtrl.Finish()
	// create a new mock object, passing the controller instance as parameter
	// for a newly created mock object it will accept any input and outpuite
	// need to define its behavior with the method expect
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}
	//
	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)
	fmt.Println(testUser.Use())
	fmt.Println()
}

func TestUser_SaySomething(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDoer := mocks.NewMockDoer(mockCtrl)
	user := User{
		Doer: mockDoer,
	}
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		expect  func()
		wantErr bool
	}{
		{
			name: "Positive test case 1",
			expect: func() {
				mockDoer.EXPECT().SaySomething("spam").Return("bad")
			},
			args:    args{num: 3},
			wantErr: false,
			want:    "bad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expect()
			if got := user.SaySomething(tt.args.num); (got != tt.want) != tt.wantErr {
				fmt.Println("gott :", got)
				t.Errorf("User.SaySomething() = %v, want %v", got, tt.want)
			}
		})
	}
}
