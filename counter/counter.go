package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Counter возвращает строку для записи в файл.
// Добавьте импорт пакета time
func Counter(count int, t time.Time) string {
    return ``
}

// Limits возвращает количество дней и запусков.
func Limits() (int, int, error) {
    // получаем имя программы
    app, err := os.Executable()
    if err != nil {
        return 0, 0, err
    }
    // получаем путь и имя текстового файла
    name := filepath.Join(filepath.Dir(app), "data.txt")
    if _, err = os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            // записываем начальные значения
            out := Counter(1, time.Now())
            // 0644 - устанавливаем разрешение на чтение и запись
            err = os.WriteFile(name, []byte(out), 0644)
            return 0, 1, err
        } 
        return 0, 0, err
    }
    return 0, 0, nil
}

func main() {
    days, counter, err := Limits()
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

    fmt.Printf("Количество дней: %d\nКоличество запусков: %d\n", days, counter)
    // устанавливаем лимит в 14 дней или 50 запусков
    if days > 14 || counter > 50 {
        fmt.Println("Запросите новую версию")
        return
    }
    fmt.Println("Программа готова к работе")
}