### Custom struct serialization

Імплементувати функцію серіалізації структру, яка використовує теги структури для своєї роботи.

Серіалізація має відповідати наступним вимогам:
* тег структури "param" дозволяє вказати назву поля структури у результаті
* пари поле->значення у результаті мають бути записані через знак "=" без пробілів (тобто `param1=x`)
* пари поле->значення розділені між собою пробілом. Приймаємо, що рядкові значення структури не можуть містити пробіл
* пари поле->значення мають бути відсортовані по назві поля у алфавітному порядку
* якщо тег "param" структури має значення "-", то це поле не має бути у результаті

