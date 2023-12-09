package main

// this is the md5 pipeline example but better i think
// the work function is simple
// gorotines for work are per item of work, goroutine are cheap.nano seconds to instgantiate
// we let the group handle the n worker for m tasks with SetLimit, just on for loop
// no select
import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

// Pipeline demonstrates the use of a Group to implement a multi-stage
// pipeline: a version of the MD5All function with bounded parallelism from
// https://blog.golang.org/pipelines.
func main() {
	m, err := MD5All(context.Background(), ".")
	if err != nil {
		log.Fatal(err)
	}

	for k, sum := range m {
		fmt.Printf("%s:\t%x\n", k, sum)
	}
}

type result struct {
	path string
	sum  [md5.Size]byte
    Err error
}

func MD5All(ctx context.Context, root string) (map[string][md5.Size]byte, error) {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(200)

	paths := make(chan string)

	g.Go(func() error {
		defer close(paths)
		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
            // TODO weg
			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	})

    // range the result
    m := make(map[string][md5.Size]byte)
	c := make(chan result)
    //go func() {
    //defer close(c)
    //    for r := range c {
    //        m[r.path] = r.sum
    //    }
    //}()
    // Why does not work?
    g.Go(func() error {
        for r := range c {
            m[r.path] = r.sum
        }
        return nil
    })

    // iterate the work 
    for path := range paths {
        g.Go(func() error {
			if err := ctx.Err(); err != nil {
				return err
			}
            
            result := workMd5(path)
			c <- result
            return nil
        })
    }

    close(c)

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
}

func workMd5(path string) result{
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return result{path, md5.Sum(data), err}
    }

    return result{path, md5.Sum(data), nil}
}

