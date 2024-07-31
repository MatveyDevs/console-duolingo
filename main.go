package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Dictionary struct {
	Words []Word
}

type Word struct {
	Word           string
	TranslatedWord string
}

func (d *Dictionary) addWord(word, translatedWord string) {
	d.Words = append(d.Words, Word{word, translatedWord})
	fmt.Println("Слово добавлено в словарь.")
}

func (d *Dictionary) showAllWords() {
	for _, word := range d.Words {
		fmt.Printf("%v -> %v\n", word.Word, word.TranslatedWord)
	}
}

func (d *Dictionary) getRandomWords() []Word {
	// Создаем копию среза слов для перемешивания
	shuffled := make([]Word, len(d.Words))
	copy(shuffled, d.Words)

	rand.Seed(time.Now().UnixNano()) // Инициализируем генератор случайных чисел
	// Алгоритм Фишера-Йетса
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled
}

func main() {
	dictionary := Dictionary{
		[]Word{
			{"tiger", "тигр"},
			{"cat", "кот"},
			{"dog", "собака"},
			{"king", "король"},
			{"queen", "королева"},
			{"queue", "очередь"},
			{"watermelon", "арбуз"},
			{"clown", "клоун"},
			{"table", "стол"},
			{"phone", "телефон"},
			{"meat", "мясо"},
			{"headphones", "наушники"},
			{"leg", "нога"},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Добавить слово")
		fmt.Println("2. Посмотреть все слова")
		fmt.Println("3. Пройти тест")
		fmt.Println("4. Выход")
		fmt.Print("Выберите действие: ")
		scanner.Scan()
		fmt.Println("")
		inputValue, _ := strconv.Atoi(scanner.Text())

		switch inputValue {
		case 1:
			fmt.Print("Введите слово: ")
			scanner.Scan()
			word := scanner.Text()
			fmt.Print("Введите перевод этого слова: ")
			scanner.Scan()
			translatedWord := scanner.Text()
			dictionary.addWord(word, translatedWord)
		case 2:
			fmt.Println("Словарь:")
			dictionary.showAllWords()
		case 3:
			shuffledDictionary := dictionary.getRandomWords()
			maxIterations := 10
			trueAnswers := 0
			count := len(shuffledDictionary)
			if count > maxIterations {
				count = maxIterations
			}
			for i, word := range shuffledDictionary[:count] {
				index := i + 1
				fmt.Printf("%d) Введите перевод слова %s: ", index, word.Word)
				scanner.Scan()
				if scanner.Text() == word.TranslatedWord {
					trueAnswers += 1
					fmt.Println("Правильно!")
				} else {
					fmt.Printf("Не правильно! Правильный ответ: %s\n", word.TranslatedWord)
				}
			}
			fmt.Printf("Молодец! Ты перевёл правильно %d слов из %d\n", trueAnswers, maxIterations)
		case 4:
			fmt.Println("Выход...")
			return

		default:
			fmt.Println("Ошибка: попробуйте снова\n")
		}

	}

}
