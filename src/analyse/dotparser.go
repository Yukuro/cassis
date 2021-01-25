package analyse

import (
	"errors"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/fatih/color"
	"io/ioutil"
)

func RequireVC(filepath string) (bool, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return false, err
	}
	graph, err := gographviz.Read(bytes)
	if err != nil {
		return false, err
	}

	for _, edge := range graph.Edges.Edges {
		res, err := requireVcFromSrcAndDstName(filepath, edge.Src, edge.Dst)
		if err != nil {
			return false, err
		}
		if res {
			return true, nil
		}

	}

	return false, nil
}

func GetAgentNameList(filepath string) []string {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		errors.New("can't read dot file")
	}
	graph, err := gographviz.Read(bytes)
	if err != nil {
		errors.New("can't parse dot file")
	}

	var agentList []string

	for _, node := range graph.Nodes.Nodes {
		agentList = append(agentList, node.Name)
	}

	return agentList
}

func Browse(filepath string) error {
	red := color.New(color.FgHiRed).PrintfFunc()
	cyan := color.New(color.FgHiCyan).PrintfFunc()
	cyanBold := color.New(color.FgHiCyan, color.Bold).PrintFunc()

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		errors.New("can't read dot file")
	}
	graph, err := gographviz.Read(bytes)
	if err != nil {
		errors.New("can't parse dot file")
	}

	fmt.Printf("\n[Analysis Results]\n")
	relation := graph.Edges.SrcToDsts
	for srcName, srcRelation := range relation {
		for dstName, _ := range srcRelation {
			cyan("[")
			fmt.Printf("%v", srcName)
			cyan("] ")
			red("-->")
			cyan(" [")
			fmt.Printf("%v", dstName)
			cyan("] ")

			res, err := requireVcFromSrcAndDstName(filepath, srcName, dstName)
			if err != nil {
				return err
			}
			if res {
				cyanBold("[VC Required]")
			}

			fmt.Println()
		}
	}
	//fmt.Println(graph)
	return nil
}

func requireVcFromSrcAndDstName(graphPath string, srcName string, dstName string) (bool, error) {
	bytes, err := ioutil.ReadFile(graphPath)
	if err != nil {
		return false, nil
	}
	graph, err := gographviz.Read(bytes)
	if err != nil {
		return false, nil
	}

	//search in Nodes index
	src_flag := false
	dst_flag := false

	for _, node := range graph.Nodes.Nodes {
		//fmt.Printf("%v == %v\n", node.Attrs["label"], "\"Issuer\"")
		if node.Name == srcName && node.Attrs["label"] == "\"Issuer\"" {
			//fmt.Println("src_flag set")
			src_flag = true
		}
	}
	for _, node := range graph.Nodes.Nodes {
		//fmt.Printf("%v == %v\n", node.Attrs["label"], "\"Holder\"")
		if node.Name == dstName && node.Attrs["label"] == "\"Holder\"" {
			//fmt.Println("dst_flag set")
			dst_flag = true
		}
	}

	if src_flag && dst_flag {
		//fmt.Println("Issuer - Holder relationship detected")
		return true, nil
	}

	return false, nil
}
