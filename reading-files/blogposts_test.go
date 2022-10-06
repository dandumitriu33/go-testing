package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"reflect"
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
			"hello world.md":  {Data: []byte("Title: hi\nDescription: x")},
			"hello-world2.md": {Data: []byte("Title: hola\nDescription: x")},
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
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(`Title: Post 1
Description: x`)},
			"hello-world2.md": {Data: []byte(`Title: Post 2
Description: x`)},
		}
		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		assertPost(t, posts[0], blogposts.Post{Title: "Post 1", Description: "x"})
	})

	t.Run("verify description is the same", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1`
			secondBody = `Title: Post 2
Description: Description 2`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
		})

	})

}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
