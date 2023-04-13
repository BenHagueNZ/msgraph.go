package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Const struct {
	Name, Value, Type string
}

type EnumType struct {
	Name        string
	Sym         string
	Members     []*EnumTypeMember
	Description string
}

type EnumTypeMember struct {
	Name        string
	Sym         string
	Value       string
	Description string
}

type EntityType struct {
	Name        string
	Sym         string
	Type        string
	Base        string
	Members     []*EntityTypeMember
	Navigations []*EntityTypeNavigation
	Description string
}

type EntityTypeMember struct {
	Name        string
	Sym         string
	Type        string
	Description string
}

type EntityTypeNavigation struct {
	Name        string
	Sym         string
	Type        string
	Description string
}

type ActionType struct {
	Name                 string
	Sym                  string
	BindingParameterType string
	Parameters           []*ActionTypeParameter
	ReturnType           string
	Description          string
}

type ActionTypeParameter struct {
	Name        string
	Sym         string
	Type        string
	Description string
}

type EntitySet struct {
	Name string
	Sym  string
	Type string
}

type Singleton struct {
	Name string
	Sym  string
	Type string
}

func (g *Generator) Generate() error {
	tmpl, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}

	inFile, err := os.Open(g.In)
	if err != nil {
		return err
	}
	defer inFile.Close()

	var o Elem
	dec := xml.NewDecoder(inFile)
	err = dec.Decode(&o)
	if err != nil {
		return err
	}
	allRequestModelMap := map[string]bool{}

	for _, schema := range o.Elems[0].Elems {
		//	alias := regexp.MustCompile("^" + schema.AttrMap()["Alias"] + "\\.")
		//	collectionAlias := regexp.MustCompile("^Collection\\(" + schema.AttrMap()["Alias"] + "\\.")
		//	collectionNamespace := "Collection(" + schema.AttrMap()["Namespace"] + "."
		namespace := "microsoft.graph."
		alias := regexp.MustCompile(`^graph\.`)
		collectionAlias := regexp.MustCompile(`^Collection\(graph\.`)
		collectionRequestAlias := regexp.MustCompile(`Request\)`)
		collectionRequestNamespace := "RequestObject)"
		requestAlias := regexp.MustCompile(`[Rr]equest$`)
		requestNamespace := "RequestObject"
		collectionNamespace := "Collection(microsoft.graph."
		namePrefix := schema.AttrMap()["Namespace"]
		fmt.Println()
		fmt.Println("Schema", namePrefix)
		if namePrefix == "microsoft.graph" {
			namePrefix = ""
		} else {
			namePrefix = strings.Replace(namePrefix, namespace, "", 1)
			runes := []rune(namePrefix)
			runes[0] = unicode.ToUpper(runes[0])
			namePrefix = string(runes)
		}

		enumTypeMap := map[string]*EnumType{}
		entityTypeMap := map[string]*EntityType{}
		actionTypeMap := map[string][]*ActionType{}
		entitySetMap := map[string]*EntitySet{}
		singletonMap := map[string]*Singleton{}
		requestModelMap := map[string]bool{}
		actionRequestBuilderMap := map[string][]string{}
		navigations := map[string]bool{}
		complete := map[string]bool{}

		for _, x := range schema.Elems {
			switch x.XMLName.Local {
			case "EnumType":
				m := x.AttrMap()
				n := m["Name"]
				runes := []rune(n)
				runes[0] = unicode.ToUpper(runes[0])
				n = namePrefix + string(runes)
				n = requestAlias.ReplaceAllString(n, requestNamespace)
				t := &EnumType{Name: n, Sym: exported(n), Description: "undocumented"}
				for _, y := range x.Elems {
					n := y.Attrs[0].Value
					// ex. Collection(String) -> StringCollection
					if strings.HasPrefix(n, colPrefix) {
						n = n[len(colPrefix):len(n)-1] + "Collection"
					}
					v := y.Attrs[1].Value
					m := &EnumTypeMember{Name: n, Sym: exported(n), Value: v, Description: "undocumented"}
					t.Members = append(t.Members, m)
				}
				enumTypeMap[n] = t
			case "EntityType", "ComplexType":
				m := x.AttrMap()
				n := m["Name"]
				if _, ok := reservedTypeTable[n]; ok {
					continue
				}
				runes := []rune(n)
				runes[0] = unicode.ToUpper(runes[0])
				n = namePrefix + string(runes)
				runes = []rune(n)
				runes[0] = unicode.ToUpper(runes[0])
				t := nsPrefix + string(runes)
				b := m["BaseType"]
				b = alias.ReplaceAllString(b, namespace)
				b = requestAlias.ReplaceAllString(b, requestNamespace)
				et := &EntityType{Name: n, Sym: exported(n), Type: t, Base: b, Description: "undocumented"}
				if strings.HasSuffix(et.Sym, "Request") {
					et.Sym += "Object"
					g.SymTypeMap[t] = et.Sym
				}
				if strings.HasSuffix(et.Type, "Request") {
					et.Type += "Object"
				}
				for _, y := range x.Elems {
					ma := y.AttrMap()
					switch y.XMLName.Local {
					case "Property":
						n := ma["Name"]
						t := alias.ReplaceAllString(ma["Type"], namespace)
						t = collectionAlias.ReplaceAllString(t, collectionNamespace)
						m := &EntityTypeMember{Name: n, Sym: exported(n), Type: t, Description: "undocumented"}
						et.Members = append(et.Members, m)
					case "NavigationProperty":
						n := ma["Name"]
						t := alias.ReplaceAllString(ma["Type"], namespace)
						t = collectionAlias.ReplaceAllString(t, collectionNamespace)
						t = collectionRequestAlias.ReplaceAllString(t, collectionRequestNamespace)
						m := &EntityTypeNavigation{Name: n, Sym: exported(n), Type: t, Description: "undocumented"}
						if strings.HasSuffix(m.Sym, "Request") {
							m.Sym += "Navigation"
						}
						et.Navigations = append(et.Navigations, m)
					}
				}
				entityTypeMap[et.Name] = et
			case "Action":
				m := x.AttrMap()
				n := m["Name"]
				runes := []rune(n)
				runes[0] = unicode.ToUpper(runes[0])
				n = namePrefix + string(runes)
				at := &ActionType{Name: n, Sym: exported(n), Description: "undocumented"}
				if strings.HasSuffix(at.Sym, "Request") {
					at.Sym += "Action"
				}
				for _, y := range x.Elems {
					ma := y.AttrMap()
					switch y.XMLName.Local {
					case "Parameter":
						n := ma["Name"]
						t := alias.ReplaceAllString(ma["Type"], namespace)
						t = collectionAlias.ReplaceAllString(t, collectionNamespace)
						t = collectionRequestAlias.ReplaceAllString(t, collectionRequestNamespace)
						if strings.HasSuffix(t, "Request") {
							t += "Object"
						}
						m := &ActionTypeParameter{Name: n, Sym: exported(n), Type: t, Description: "undocumented"}
						at.Parameters = append(at.Parameters, m)
					case "ReturnType":
						at.ReturnType = alias.ReplaceAllString(ma["Type"], namespace)
						at.ReturnType = collectionAlias.ReplaceAllString(at.ReturnType, collectionNamespace)
					}
				}
				at.BindingParameterType = at.Parameters[0].Type
				at.Parameters = at.Parameters[1:]
				actionTypeMap[at.BindingParameterType] = append(actionTypeMap[at.BindingParameterType], at)
			case "EntityContainer":
				for _, y := range x.Elems {
					ma := y.AttrMap()
					switch y.XMLName.Local {
					case "EntitySet":
						s := &EntitySet{
							Name: ma["Name"],
							Sym:  exported(ma["Name"]),
							Type: alias.ReplaceAllString(ma["EntityType"], namespace),
						}
						entitySetMap[s.Name] = s
					case "Singleton":
						s := &Singleton{
							Name: ma["Name"],
							Sym:  exported(ma["Name"]),
							Type: alias.ReplaceAllString(ma["Type"], namespace),
						}
						singletonMap[s.Name] = s
					}
				}
			case "Annotations":
				mas := x.AttrMap()
				target, _ := stripNSPrefix(mas["Target"])
				targetList := strings.Split(target, "/")
				targetMember := ""
				if len(targetList) > 1 {
					target = alias.ReplaceAllString(targetList[0], namespace)
					target = requestAlias.ReplaceAllString(target, requestNamespace)
					targetMember = alias.ReplaceAllString(targetList[1], namespace)
					targetMember = requestAlias.ReplaceAllString(targetMember, requestNamespace)
				}
				target = alias.ReplaceAllString(target, namespace)
				target = requestAlias.ReplaceAllString(target, requestNamespace)

				targetMember = alias.ReplaceAllString(targetMember, namespace)
				targetMember = requestAlias.ReplaceAllString(targetMember, requestNamespace)
				for _, y := range x.Elems {
					switch y.XMLName.Local {
					case "Annotation":
						ma := y.AttrMap()
						term := ma["Term"]
						str := ma["String"]
						if term == "Org.OData.Core.V1.Description" {
							if e, ok := entityTypeMap[target]; ok {
								if targetMember == "" {
									e.Description = str
								} else {
									for _, mem := range e.Members {
										if mem.Name == targetMember {
											mem.Description = str
										}
									}
								}
							}
						}
					}
				}
			}
		}
		//	continue
		//	for x := range entityTypeMap {
		//		fmt.Println(x)
		//	}
		//	os.Exit(0)

		var out io.WriteCloser

		keys := []string{}
		for x := range enumTypeMap {
			keys = append(keys, x)
		}
		sort.Strings(keys)
		for _, key := range keys {
			x := enumTypeMap[key]
			out, err = g.Create("Enum", x.Sym)
			g.X = x
			if err != nil {
				return err
			}
			err := tmpl.ExecuteTemplate(out, "enum.go.tmpl", g)
			if err != nil {
				return err
			}
			out.Close()
		}

		keys = nil
		for x := range entityTypeMap {
			keys = append(keys, x)
		}
		sort.Strings(keys)

		for _, key := range keys {
			x := entityTypeMap[key]
			x.Type = alias.ReplaceAllString(x.Type, namespace)
			x.Base = alias.ReplaceAllString(x.Base, namespace)

			out, err = g.Create("Model", x.Sym)
			if err != nil {
				return err
			}
			g.X = x

			err := tmpl.ExecuteTemplate(out, "model.go.tmpl", g)
			if err != nil {
				return err
			}
			out.Close()
		}

		keys = nil
		for x := range actionTypeMap {
			keys = append(keys, x)
		}
		sort.Strings(keys)
		for _, a := range keys {
			a = collectionAlias.ReplaceAllString(a, collectionNamespace)
			a = alias.ReplaceAllString(a, namespace)
			if _, ok := reservedTypeTable[a]; ok {
				continue
			}
			x := actionTypeMap[a]
			out, err = g.Create("Action", g.SymBaseType(a))
			if err != nil {
				return err
			}
			if !allRequestModelMap[g.SymBaseType(a)] {
				requestModelMap[g.SymBaseType(a)] = true
				allRequestModelMap[g.SymBaseType(a)] = true
			}
			for _, y := range x {
				g.X = y
				err := tmpl.ExecuteTemplate(out, "action.go.tmpl", g)
				if err != nil {
					return err
				}
			}
			out.Close()
		}
		for _, x := range entityTypeMap {
			//	if len(x.Navigations) == 0 {
			//		fmt.Println("a")
			//		fmt.Println(x.Sym)
			//		continue
			//	}
			if !allRequestModelMap[x.Sym] {
				requestModelMap[x.Sym] = true
				allRequestModelMap[x.Sym] = true
			}
			for _, y := range x.Navigations {
				if !allRequestModelMap[g.SymBaseType(y.Type)] {
					requestModelMap[g.SymBaseType(y.Type)] = true
					allRequestModelMap[g.SymBaseType(y.Type)] = true
				}
			}
		}

		for _, x := range entitySetMap {
			if !allRequestModelMap[g.SymBaseType(x.Type)] {
				requestModelMap[g.SymBaseType(x.Type)] = true
				allRequestModelMap[g.SymBaseType(x.Type)] = true
			}
		}
		for _, x := range singletonMap {
			if !allRequestModelMap[g.SymBaseType(x.Type)] {
				requestModelMap[g.SymBaseType(x.Type)] = true
				allRequestModelMap[g.SymBaseType(x.Type)] = true
			}
		}

		keys = nil
		for x := range requestModelMap {
			keys = append(keys, x)

		}
		sort.Strings(keys)
		for _, x := range keys {
			out, err = g.Create("Request", x)
			if err != nil {
				return err
			}

			g.X = x
			err := tmpl.ExecuteTemplate(out, "request_model.go.tmpl", g)
			if err != nil {
				return err
			}
			out.Close()
		}

		out, err = g.Create("GraphService", "")
		if err != nil {
			return err
		}
		keys = nil
		for x := range entitySetMap {
			keys = append(keys, x)
		}
		sort.Strings(keys)
		for _, key := range keys {
			g.X = &EntityType{Name: "GraphService", Sym: "GraphService"}
			g.Y = entitySetMap[key]
			err := tmpl.ExecuteTemplate(out, "request_collection_navigation.go.tmpl", g)
			if err != nil {
				return err
			}
		}
		keys = nil
		for x := range singletonMap {
			keys = append(keys, x)
		}
		sort.Strings(keys)
		for _, key := range keys {
			g.X = &EntityType{Name: "GraphService", Sym: "GraphService"}
			g.Y = singletonMap[key]
			err := tmpl.ExecuteTemplate(out, "request_navigation.go.tmpl", g)
			if err != nil {
				return err
			}
		}
		out.Close()

		keys = nil
		for x := range entityTypeMap {
			keys = append(keys, x)
		}

		sort.Strings(keys)
		for _, key := range keys {
			x := entityTypeMap[key]
			if !contains(actionRequestBuilderMap[strings.ToLower(x.Type)], x.Sym) {
				actionRequestBuilderMap[strings.ToLower(x.Type)] = append(actionRequestBuilderMap[strings.ToLower(x.Type)], x.Sym)
			}
			if len(x.Navigations) == 0 {
				continue
			}
			out, err = g.Create("Action", x.Sym)
			if err != nil {
				return err
			}
			g.X = x
			sort.Slice(x.Navigations, func(i, j int) bool { return x.Navigations[i].Name < x.Navigations[j].Name })
			for _, y := range x.Navigations {
				navigations[y.Type] = true
				g.Y = y
				if isCollectionType(y.Type) {
					if !contains(actionRequestBuilderMap[strings.ToLower(y.Type)], x.Sym+y.Sym+"Collection") {
						actionRequestBuilderMap[strings.ToLower(y.Type)] = append(actionRequestBuilderMap[strings.ToLower(y.Type)], x.Sym+y.Sym+"Collection")
					}
					err := tmpl.ExecuteTemplate(out, "request_collection_navigation.go.tmpl", g)
					if err != nil {
						return err
					}
				} else {
					err := tmpl.ExecuteTemplate(out, "request_navigation.go.tmpl", g)
					if err != nil {
						return err
					}
				}
			}
			out.Close()
		}
		fmt.Println(actionRequestBuilderMap)
		for _, key := range keys {
			x := entityTypeMap[key]
			out, err = g.Create("Action", x.Sym)
			if err != nil {
				return err
			}
			if !navigations[x.Type] && !navigations["Collection("+x.Type+")"] && (x.Base != "") {
				fmt.Println("a")
				fmt.Println(x.Type)
				fmt.Println(x.Base)
				fmt.Println(x.Sym)
				fmt.Println(key)
				baseType := g.SymFromType(x.Base)
				if strings.HasSuffix(baseType, "RequestObject") {
					baseType = strings.TrimSuffix(baseType, "Object")
				}
				if !complete[x.Sym] {
					if navigations["Collection("+x.Base+")"] {
						g.Y = x
						y := entityTypeMap[baseType]
						//y.Sym = x.Sym
						g.X = y
						complete[y.Sym] = true
						fmt.Println("c")
						fmt.Println(y.Sym)
						fmt.Println("Collection(" + y.Type + ")")
						fmt.Println(actionRequestBuilderMap[strings.ToLower("Collection("+y.Type+")")])
						val, ok := actionRequestBuilderMap[strings.ToLower("Collection("+y.Type+")")]
						if ok {
							y.Sym = val[0]
						}
						x.Name = ""
						if !contains(actionRequestBuilderMap[strings.ToLower(x.Type)], x.Sym+"Collection") {
							actionRequestBuilderMap[x.Type] = append(actionRequestBuilderMap[strings.ToLower(x.Type)], x.Sym+"Collection")
							fmt.Println("Contains1")
						}
						err := tmpl.ExecuteTemplate(out, "request_collection_navigation.go.tmpl", g)
						if err != nil {
							return err
						}
					} else {
						if navigations[x.Base] {
							g.X = x
							y := entityTypeMap[baseType]
							//y.Sym = x.Sym
							g.Y = y
							complete[y.Sym] = true
							fmt.Println("b")
							fmt.Println(y.Sym)
							//val, ok := actionRequestBuilderMap[strings.ToLower(y.Type)]
							//if ok {
							//	x.Sym = val[0]
							//}
							if !contains(actionRequestBuilderMap[strings.ToLower(x.Type)], x.Sym) {
								actionRequestBuilderMap[strings.ToLower(x.Type)] = append(actionRequestBuilderMap[strings.ToLower(x.Type)], x.Sym)
								fmt.Println("Contains2")
							}
							err := tmpl.ExecuteTemplate(out, "request_navigation.go.tmpl", g)
							if err != nil {
								return err
							}
						}
					}
				}
			}
			out.Close()
		}
		keys = nil
		for x := range actionTypeMap {
			keys = append(keys, x)
		}
		sort.Strings(keys)
		for _, a := range keys {

			x := actionTypeMap[a]
			if _, ok := reservedTypeTable[a]; ok {
				continue
			}
			for _, y := range x {
				out, err = g.Create("Request", g.SymBaseType(a)+y.Sym)
				if err != nil {
					return err
				}
				g.Y = y
				if b, ok := actionRequestBuilderMap[strings.ToLower(a)]; ok {
					g.X = b

					if y.ReturnType == "" {
						err = tmpl.ExecuteTemplate(out, "request_action_void.go.tmpl", g)
					} else if isCollectionType(y.ReturnType) {
						err = tmpl.ExecuteTemplate(out, "request_action_collection.go.tmpl", g)
					} else {
						err = tmpl.ExecuteTemplate(out, "request_action_single.go.tmpl", g)
					}
					if err != nil {
						return err
					}
				}
				out.Close()
			}
		}
	}
	// Copy all *.go files without template processing
	// Copy everything below the first "// BEGIN" line to the output
	paths, err := filepath.Glob("templates/*.go")
	if err != nil {
		return err
	}
	for _, path := range paths {
		err := func() error {
			inFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer inFile.Close()
			outFile, err := g.Create(filepath.Base(path), "")
			if err != nil {
				return err
			}
			defer outFile.Close()
			scanner := bufio.NewScanner(inFile)
			enabled := false
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "// BEGIN") {
					enabled = true
					continue
				}
				if enabled {
					fmt.Fprintln(outFile, line)
				}
			}
			return scanner.Err()
		}()
		if err != nil {
			return err
		}
	}
	// Copy some *.go.tmpl files with template processing
	for _, path := range []string{"const.go.tmpl"} {
		out, err := g.Create(path[:len(path)-5], "")
		if err != nil {
			return err
		}
		err = tmpl.ExecuteTemplate(out, path, g)
		if err != nil {
			return err
		}
		out.Close()
	}
	return nil
}
