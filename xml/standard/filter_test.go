package xml

import (
	"bytes"
	"testing"

	"github.com/admpub/xencoding/filter"
)

func TestFilter(t *testing.T) {
	type Post struct {
		ID    int64
		Title string
		Body  string
	}
	type Profile struct {
		Name string
		Age  int
		Post Post
	}
	type User struct {
		ID       int64
		Password string
		Profile  Profile
	}
	u := User{
		ID:       1,
		Password: "password",
		Profile: Profile{
			Name: "name",
			Age:  20,
			Post: Post{
				ID:    1,
				Title: "title",
				Body:  "body",
			},
		},
	}
	r, err := MarshalFilter(u, filter.Exclude("Password"))
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal([]byte(`<User><ID>1</ID><Profile><Name>name</Name><Age>20</Age><Post><ID>1</ID><Title>title</Title><Body>body</Body></Post></Profile></User>`), r) {
		t.Log(string(r))
		t.Error("MarshalFilter failed")
	}

	r, err = MarshalSelector(u, filter.Include("ID", "Password", "Profile.*"))
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal([]byte(`<User><ID>1</ID><Password>password</Password><Profile><Name>name</Name><Age>20</Age><Post><ID>1</ID><Title>title</Title><Body>body</Body></Post></Profile></User>`), r) {
		t.Log(string(r))
		t.Error("MarshalSelector failed")
	}

	r, err = MarshalWithOption(u,
		OptionFilter(filter.Exclude("Password", "Profile.Post.Body")),
		OptionSelector(filter.Include("ID", "Password", "Profile.*")),
	)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal([]byte(`<User><ID>1</ID><Profile><Name>name</Name><Age>20</Age><Post><ID>1</ID><Title>title</Title></Post></Profile></User>`), r) {
		t.Log(string(r))
		t.Error("MarshalWithOption failed")
	}
}
