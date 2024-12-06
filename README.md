# Часть сервиса аутентификации

### Перед использованием
### Необходимо создать и настроить .env файл
```
PORT=       --порт базы данных
DB=         --данные для подключения базы данных
SECRET=     --секретная фраза для создания подписи JWT
EMAIL=      --адрес почты для отправки предупреждений
MAILPASS=   --пароль для использования почты сторонними сервисами
```
Примечание: 
- DB должен выглядеть в формате "\<host\> user=\<username\> password=\<password\> dbname=\<dbname\> port=\<port\> sslmode=\<sslmode\>"
- MAILPASS это не пароль от почты, а специальны пароль для использования сторонними сервисами ([*Подробнее для mail.ru*](https://help.mail.ru/mail/security/protection/external/))