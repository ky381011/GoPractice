package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// Terraform ファイルのトップレベルスキーマ
// ブロックタイプとラベル数を定義する
var tfSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: "provider", LabelNames: []string{"name"}},
		{Type: "variable", LabelNames: []string{"name"}},
		{Type: "resource", LabelNames: []string{"type", "name"}},
	},
}

func main() {
	parser := hclparse.NewParser()

	// .tf ファイルは HCL2 形式なので ParseHCLFile で読み込める
	file, diags := parser.ParseHCLFile("main.tf")
	if diags.HasErrors() {
		fmt.Println("Parse error:", diags.Error())
		return
	}

	// スキーマに従ってブロックを抽出
	content, diags := file.Body.Content(tfSchema)
	if diags.HasErrors() {
		fmt.Println("Content error:", diags.Error())
		return
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "provider":
			fmt.Printf("[provider] %s\n", block.Labels[0])
		case "variable":
			fmt.Printf("[variable] %s\n", block.Labels[0])
		case "resource":
			fmt.Printf("[resource] %s \"%s\"\n", block.Labels[0], block.Labels[1])
		}
		printAttrs(block.Body, file.Bytes)
		fmt.Println()
	}
}

// JustAttributes でブロック内の全属性を取得して表示する
// ネストブロックが含まれる場合は PartialContent にフォールバック
// 値はソースバイト列から Range で直接スライスするため cty 不要
func printAttrs(body hcl.Body, src []byte) {
	attrs, diags := body.JustAttributes()
	if diags.HasErrors() {
		// ネストしたブロックがあると JustAttributes は失敗するので
		// PartialContent で属性だけ取り出す
		partial, _, _ := body.PartialContent(&hcl.BodySchema{})
		attrs = partial.Attributes
	}

	for name, attr := range attrs {
		r := attr.Expr.Range()
		val := string(src[r.Start.Byte:r.End.Byte])
		fmt.Printf("  %-20s = %s\n", name, val)
	}
}
