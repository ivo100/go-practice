package data

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"strings"
)

/*
https://www.freecodecamp.org/news/exploratory-data-analysis-in-go-with-gota/
*/

// Run applies a function (mean) to each column
func Run0() {
	df := dataframe.New(
		series.New([]int{10, 10, 10, 10, 10}, series.Int, "score"),
	)

	mean := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		return series.Floats(sum / float64(len(floats)))
	}

	meanScore := df.Capply(mean)

	fmt.Println(meanScore)
}
func Run() {
	df := SomeDf()

	description := df.Describe()

	fmt.Println(description)

	fmt.Println(df)

	// This selects  rows
	rows := df.Subset([]int{0, 1})
	fmt.Println(rows)
	// This selects  columns
	cols := df.Select([]int{0, 3, 1})
	fmt.Println(cols)

	cols = df.Select([]string{"Age", "Name"})
	fmt.Println(cols)
	// UPDATE
	updatedDf := df.Set(
		[]int{0, 2},
		dataframe.LoadRecords(
			[][]string{
				[]string{"Jenny", "23", "Purple", "2.2"},
				[]string{"Jesse", "34", "Indigo", "3.5"},
				[]string{"Peter", "33", "Violet", "3.3"},
			},
		),
	)
	fmt.Println(updatedDf)
	// FILTER
	filtered := df.Filter(
		dataframe.F{Colname: "Age", Comparator: series.Greater, Comparando: 20},
	)
	fmt.Println(filtered)

	//SORT
	sorted := df.Arrange(
		dataframe.Sort("Age"),
	)
	fmt.Println(sorted)

	group := df.GroupBy("Age")
	fmt.Println(group)

}

func SomeDf() dataframe.DataFrame {
	csvString := `
Name,Age,Favorite Color,Height(ft)
John,18,Red,6.7
Bob,20,Blue,5.11
Mary,20,Blue,5.7
`

	csvDf := dataframe.ReadCSV(strings.NewReader(csvString))
	return csvDf
}
func Run2() {
	/*
	   	// Create a new dataframe.
	   	df := dataframe.New(
	   		series.New([]string{"a", "b", "c"}, series.String, "letter"),
	   		series.New([]int{1, 2, 3}, series.Int, "number"),
	   	)
	   	fmt.Printf("%v\n", df)

	   	type Dog struct {
	   		Name       string
	   		Color      string
	   		Height     int
	   		Vaccinated bool
	   	}

	   	dogs := []Dog{
	   		{"Buster", "Black", 56, false},
	   		{"Jake", "White", 61, false},
	   		{"Bingo", "Brown", 50, true},
	   		{"Gray", "Cream", 68, false},
	   	}

	   	dogsDf := dataframe.LoadStructs(dogs)

	   	fmt.Println(dogsDf)

	   	fmt.Println(dogsDf.Dims())
	   	fmt.Println(dogsDf.Types())
	   	fmt.Println(dogsDf.Names())
	   	fmt.Println(dogsDf.Nrow())
	   	fmt.Println(dogsDf.Ncol())

	   	csvString := `
	   Name,Age,Favorite Color,Height(ft)
	   John,44,Red,6.7
	   Mary,40,Blue,5.7`

	   	csvDf := dataframe.ReadCSV(strings.NewReader(csvString))
	   	fmt.Println(csvDf)

	   	// This selects the first two rows of the DataFrame
	   	rows := df.Subset([]int{0, 2})
	   	fmt.Println(rows)

	   	/*
	   			jsonString := `[
	   		  {
	   		    "Name": "John",
	   		    "Age": 44,
	   		    "Favorite Color": "Red",
	   		    "Height(ft)": 6.7
	   		  },
	   		  {
	   		    "Name": "Mary",
	   		    "Age": 40,
	   		    "Favorite Color": "Blue",
	   		    "Height(ft)": 5.7
	   		  }
	   		]`

	   			jsonDf := dataframe.ReadJSON(strings.NewReader(jsonString))
	   			fmt.Println(jsonDf)
	*/
}
