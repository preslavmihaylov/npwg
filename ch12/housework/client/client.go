package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	pbgen "housework/idl"
)

var addr string

func init() {
	flag.StringVar(&addr, "address", "localhost:34443",
		"server address")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			`Usage: %s [flags] [add chore, ...|complete #]
    add         add comma-separated chores
    complete    complete designated chore
Flags:
`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func list(ctx context.Context, client pbgen.RobotMaidClient) error {
	chores, err := client.List(ctx, new(pbgen.Empty))
	if err != nil {
		return err
	}

	if len(chores.Chores) == 0 {
		fmt.Println("You have nothing to do!")
		return nil
	}

	fmt.Println("#\t[X]\tDescription")
	for i, chore := range chores.Chores {
		c := " "
		if chore.Complete {
			c = "X"
		}
		fmt.Printf("%d\t[%s]\t%s\n", i+1, c, chore.Description)
	}

	return nil
}

func add(ctx context.Context, client pbgen.RobotMaidClient, s string) error {
	chores := new(pbgen.Chores)

	for _, chore := range strings.Split(s, ",") {
		if desc := strings.TrimSpace(chore); desc != "" {
			chores.Chores = append(chores.Chores, &pbgen.Chore{
				Description: desc,
			})
		}
	}

	_, err := client.Add(ctx, chores)

	return err
}

func complete(ctx context.Context, client pbgen.RobotMaidClient, s string) error {
	i, err := strconv.Atoi(s)
	if err == nil {
		_, err = client.Complete(ctx,
			&pbgen.CompleteRequest{ChoreNumber: int32(i)})
	}

	return err
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	rosie := pbgen.NewRobotMaidClient(conn)
	ctx := context.Background()

	switch strings.ToLower(flag.Arg(0)) {
	case "add":
		err = add(ctx, rosie, strings.Join(flag.Args()[1:], " "))
	case "complete":
		err = complete(ctx, rosie, flag.Arg(1))
	}

	if err != nil {
		log.Fatal(err)
	}

	err = list(ctx, rosie)
	if err != nil {
		log.Fatal(err)
	}
}
