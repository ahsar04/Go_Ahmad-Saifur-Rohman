package main

import "fmt"

type Student struct {
    name  []string
    score []float64
}

func (s *Student) min() (string, float64) {
    min := s.score[0]
    minName := s.name[0]
    for i, score := range s.score {
        if score < min {
            min = score
            minName = s.name[i]
        }
    }
    return minName, min
}

func (s *Student) max() (string, float64) {
    max := s.score[0]
    maxName := s.name[0]
    for i, score := range s.score {
        if score > max {
            max = score
            maxName = s.name[i]
        }
    }
    return maxName, max
}

func (s *Student) avg() float64 {
    var sum float64
    for _, score := range s.score {
        sum += score
    }
    return sum / float64(len(s.score))
}

func main() {
    var students [5]Student
	
	fmt.Println("Input: ")
    for i := 0; i < 5; i++ {
        var nameData string
        var scoreData float64
        for {
			fmt.Printf("Input %d Student's Name ", i+1)
            fmt.Scanln(&nameData)

            fmt.Printf("Input %d Student's Score ", i+1)
            _, err := fmt.Scanln(&scoreData)
            if err == nil {
                break
            }
            fmt.Println("Input salah, masukkan angka")
        }

        students[i] = Student{name: []string{nameData}, score: []float64{scoreData}}
    }
	// var  min, max, sum float64
	// min=students[0].score[0]
    // for _, student := range students {
	// 	for i := 0; i < len(student.score); i++ {
	// 		if min > student.score[i] {
	// 			min = student.score[i]
	// 		}
	// 		if max < student.score[i] {
	// 			max = student.score[i]
	// 		}
	// 		sum = sum + student.score[i]
	// 	}
		// print Score
		// for _, score := range student.score {
		// 	fmt.Printf("%v\n",score)
		// }
    // } 
	// fmt.Println("Average Score ",sum/5 )
	// fmt.Println("Min Score of Student (",min,")" )
	// fmt.Println("Max Score of Student (",max,")" )

	// fmt.Print(students)
	fmt.Println()
	fmt.Println("Output: ")
	var totalScore float64
    for _, student := range students {
        totalScore += student.avg()
    }
    fmt.Println("Average Score ", totalScore/5)

	var minName, maxName string
    var min float64 = students[0].score[0]
    var max float64 = students[0].score[0]
    for _, student := range students {
        name, score := student.min()
        if score < min {
            minName = name
            min = score
        }
        name, score = student.max()
        if score > max {
            maxName = name
            max = score
        }
    }
	
    fmt.Printf("Min Score of Students  %s (%v)\n", minName, min)
    fmt.Printf("Max Score of Students  %s (%v)\n", maxName, max)
}
