/*
 * Minio tool to be used with Pivotal "cf" tool to generate the
 * config for create-service command.
 */

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/minio/cli"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "Minio CF Tool"
	app.Author = "https://minio.io"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "access-key",
		},
		cli.StringFlag{
			Name: "secret-key",
		},
		cli.StringFlag{
			Name: "gcs",
		},
		cli.BoolFlag{
			Name: "azure",
		},
		cli.BoolFlag{
			Name: "s3",
		},
		cli.StringFlag{
			Name: "subdomain",
		},
	}
	app.CustomAppHelpTemplate = `NAME:
  {{.HelpName}} - {{.Usage}}

USAGE:
  {{.HelpName}} [FLAGS]

FLAGS:
  {{range .VisibleFlags}}{{.}}
  {{end}}

EXAMPLES:
  1. Generate config file for Minio Server instance
     $ {{.HelpName}} --access-key minio --secret-key minio123 > minio-server-pcf.conf

  2. Generate config file for Minio GCS gateway instance
     $ {{.HelpName}} --access-key minio --secret-key minio123 --gcs /path/to/credentials.json > minio-gcs-pcf.conf

  3. Generate config file for Minio Azure gateway instance
     $ {{.HelpName}} --access-key azureaccountname --secret-key azureaccountkey > minio-azure-pcf.conf

`

	app.Action = func(ctx *cli.Context) {
		accessKey := ctx.GlobalString("access-key")
		secretKey := ctx.GlobalString("secret-key")
		if accessKey == "" || secretKey == "" {
			cli.ShowAppHelpAndExit(ctx, 1)
		}
		m := make(map[string]string)
		m["accesskey"] = accessKey
		m["secretkey"] = secretKey
		subdomain := ctx.GlobalString("subdomain")
		if subdomain != "" {
			m["subdomain"] = subdomain
		}
		if ctx.GlobalBool("azure") {
			m["gateway"] = "azure"
		}
		if ctx.GlobalBool("s3") {
			m["gateway"] = "s3"
		}
		gcsPath := ctx.GlobalString("gcs")
		if gcsPath != "" {
			m["gateway"] = "gcs"
			cleanPath := filepath.Clean(gcsPath)
			b, err := os.ReadFile(cleanPath)
			if err != nil {
				log.Fatal(err)
			}
			m["googlecredentials"] = base64.StdEncoding.EncodeToString(b)
		}
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		writer := ctx.App.Writer
		if writer == nil {
			writer = os.Stdout
		}
		fmt.Fprintf(writer, "%s", string(b))
	}
	return app
}

func main() {
	app := newApp()
	app.Run(os.Args)
}
