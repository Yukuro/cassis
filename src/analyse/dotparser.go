package analyse

import (
	"errors"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/fatih/color"
	"io/ioutil"
)

func RequireVC(filepath string) bool {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		errors.New("can't read dot file")
	}
	graph, err := gographviz.Read(bytes)
	if err != nil {
		errors.New("can't parse dot file")
	}

	for _, edge := range graph.Edges.Edges {
		//search in Nodes index
		src_flag := false
		dst_flag := false

		for _, node := range graph.Nodes.Nodes {
			//fmt.Printf("%v == %v\n", node.Attrs["label"], "\"Issuer\"")
			if edge.Src == node.Name && node.Attrs["label"] == "\"Issuer\"" {
				//fmt.Println("src_flag set")
				src_flag = true
			}
		}
		for _, node := range graph.Nodes.Nodes {
			//fmt.Printf("%v == %v\n", node.Attrs["label"], "\"Holder\"")
			if edge.Dst == node.Name && node.Attrs["label"] == "\"Holder\"" {
				//fmt.Println("dst_flag set")
				dst_flag = true
			}
		}

		if src_flag && dst_flag {
			//fmt.Println("Issuer - Holder relationship detected")
			return true
		}

	}

	return false
}

func GetAgentNameList(filepath string) []string{
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		errors.New("can't read dot file")
	}
	graph, err := gographviz.Read(bytes)
	if err != nil {
		errors.New("can't parse dot file")
	}

	var agentList []string

	for _, node := range graph.Nodes.Nodes{
		agentList = append(agentList, node.Name)
	}

	return agentList
}

func Browse(filepath string) {
	red := color.New(color.FgHiRed).PrintfFunc()
	cyan := color.New(color.FgHiCyan).PrintfFunc()

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		errors.New("can't read dot file")
	}
	graph, err := gographviz.Read(bytes)
	if err != nil {
		errors.New("can't parse dot file")
	}

	fmt.Println("[relation]")
	relation := graph.Edges.SrcToDsts
	for srcName, srcRelation := range relation {
		for dstName, _ := range srcRelation {
			cyan("[")
			fmt.Printf("%v", srcName)
			cyan("] ")
			red("-->")
			cyan(" [")
			fmt.Printf("%v", dstName)
			cyan("]\n")
		}
	}
	//fmt.Println(graph)
}




