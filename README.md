# Регулярные выражения

## Простое совпадение текста

Большинство букв и символов соответствуют сами себе.
Символы, которые не соответствуют сами себе (метасимволы): `. ^ $ * + ? { [ ] \ | ( )`
Если в регулярку написать часть текста (просто буквами), то совпадение сработает, если этот кусок пристутствует в проверяемом тексте.

```
"when bob will meet us" =~ /bob/ -> true
"when we'll meet robert" =~ /bob/ -> false
```

*Регулярка матчит любой кусок строки.*

Первые метасимволы:
`^` — Начало строки.
`$` — Конец строки.

```
"bob" =~ /^bob$/ —> true
"bob roth" =~ /^bob$/ —> false
"bob roth" =~ /^bob/ —> true
"who is bob?" =~ /^bob/ —> false
"this is bob" =~ /bob$/ —> true
"this is bob's sister" =~ /bob$/ —> false
```

## Классы символов

`[abc]` — соответствие любой из трех букв a, b или c.
`[a-z]` — соответствие любой строчной английской букве (от a до z).
`[0-9a-z]` — соответствие любой строчной английской букве (от a до z) или числу от 0 до 9.
`[^abc]` — метасимвол `^` в начале класса означает "всё, кроме abc".

```
"order=5" =~ /order=[0-9]/ —> true
"service_letter=f" =~ /service_letter=[a-z]/ —> true
"hexcode=a7bf" =~ /hexcode=[0-9a-f][0-9a-f][0-9a-f][0-9a-f]/ —> true
"not_digits=5" =~ /not_digits=[^0-9]/ —> false
```

Метасимволы теряют свою силу внутри классов (кроме `\ ] ^ -`).
Потому что:
- `\` для экранирования
- `]` для обозначения конца класса
- `^` отрицание (здесь он не означает начало строки)
- `-` диапазон
```
"a*b+c" =~ /a[+-*]b[+-*]c/ -> true
```

### Встроенные классы

Некоторые из них:
`\d` — любая цифра. `[0-9]`
`\w` — латинская буква, цифра или подчеркивание. `[a-zA-Z0-9_]`
`\s` — пробельный символ. `[\r\n\t\f\v ]`

## Квантификаторы

Определяют количество повторений.

