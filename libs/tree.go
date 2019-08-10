package libs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Tree(root, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}

	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", root, err)
	}

	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}

	sign := "├──"
	for i, name := range names {
		add := "│  "
		if i == len(names)-1 {
			sign = "└──"
			add = "   "
		}

		fmt.Printf(indent + sign)

		if err := Tree(filepath.Join(root, name), indent+add); err != nil {
			return err
		}
	}
	return nil
}
