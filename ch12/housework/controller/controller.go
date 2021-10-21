package controller

import (
	"fmt"
	"housework/storage"
	"os"
	"strconv"
	"strings"
)

func New(s storage.Storage, dbFile string) *Controller {
	return &Controller{
		storage: s,
		dbFile:  dbFile,
	}
}

type Controller struct {
	storage storage.Storage
	dbFile  string
}

func (c *Controller) Load() ([]*storage.Chore, error) {
	if _, err := os.Stat(c.dbFile); os.IsNotExist(err) {
		return []*storage.Chore{}, nil
	}

	df, err := os.Open(c.dbFile)
	if err != nil {
		return nil, err
	}
	defer df.Close()

	// TODO:
	return c.storage.Load(df)
}

func (c *Controller) Flush(chores []*storage.Chore) error {
	df, err := os.Create(c.dbFile)
	if err != nil {
		return err
	}
	defer df.Close()

	return c.storage.Flush(df, chores)
}

func (c *Controller) List() error {
	chores, err := c.Load()
	if err != nil {
		return err
	}

	if len(chores) == 0 {
		fmt.Println("You're all caught up!")
		return nil
	}

	fmt.Println("#\t[X]\tDescription")
	for i, chore := range chores {
		c := " "
		if chore.Complete {
			c = "X"
		}
		fmt.Printf("%d\t[%s]\t%s\n", i+1, c, chore.Description)
	}

	return nil
}

func (c *Controller) Add(s string) error {
	chores, err := c.Load()
	if err != nil {
		return err
	}

	for _, chore := range strings.Split(s, ",") {
		if desc := strings.TrimSpace(chore); desc != "" {
			chores = append(chores, &storage.Chore{
				Description: desc,
			})
		}
	}

	return c.Flush(chores)
}

func (c *Controller) Complete(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	chores, err := c.Load()
	if err != nil {
		return err
	}

	if i < 1 || i > len(chores) {
		return fmt.Errorf("chore %d not found", i)
	}

	chores[i-1].Complete = true
	return c.Flush(chores)
}
