package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {

	t.Run("simple count of the number of files in folder", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("folder read should fail", func(t *testing.T) {

		errorMessage := "oh no, I always fail"

		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err.Error() != errorMessage {
			t.Errorf(`did not get an error or it is not "%s"`, errorMessage)
		}
	})

	t.Run("verify title is the same", func(t *testing.T) {

	})

}
