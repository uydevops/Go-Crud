# Go CRUD Application

Bu proje, Go dili kullanılarak geliştirilmiş basit bir CRUD (Create, Read, Update, Delete) uygulamasıdır. SQLite veritabanını kullanarak kullanıcıları yönetir ve HTTP sunucusu üzerinden RESTful API sağlar.

## Özellikler

- Kullanıcıları listeleme
- Kullanıcı ekleme
- Kullanıcı güncelleme
- Kullanıcı silme

## Başlarken

Bu proje, Go dilinde geliştirilmiştir. Başlamak için aşağıdaki adımları takip edin:

### Gereksinimler

- Go (1.16 veya üzeri)
- SQLite3

### Kurulum

1. **Depoyu klonlayın:**

    ```sh
    git clone https://github.com/kullaniciadi/go-crud-app.git
    cd go-crud-app
    ```

2. **Gerekli Go modüllerini yükleyin:**

    ```sh
    go mod tidy
    ```

3. **Çevresel değişkenleri ayarlayın (isteğe bağlı):**

    `.env` dosyası oluşturabilir ve gerekli çevresel değişkenleri tanımlayabilirsiniz:

    ```
    DATABASE_DSN=file:test.db
    PORT=:8080
    ```

### Kullanım

1. **Sunucuyu başlatın:**

    ```sh
    go run cmd/server/main.go
    ```

    Sunucu varsayılan olarak `http://localhost:8080` adresinde çalışacaktır.

2. **API Uç Noktaları:**

    - **Kullanıcıları listele:** `GET /users`
    - **Kullanıcı oluştur:** `POST /users/create` (JSON formatında `name` ve `email` içermelidir)
    - **Kullanıcı güncelle:** `PUT /users/update` (JSON formatında `id`, `name` ve `email` içermelidir)
    - **Kullanıcı sil:** `DELETE /users/delete?id=<id>` (Query parametre olarak `id` içermelidir)

### Örnek İstekler

- **Kullanıcı oluşturma:**

    ```sh
    curl -X POST http://localhost:8080/users/create \
    -H "Content-Type: application/json" \
    -d '{"name": "John Doe", "email": "john.doe@example.com"}'
    ```

- **Kullanıcı güncelleme:**

    ```sh
    curl -X PUT http://localhost:8080/users/update \
    -H "Content-Type: application/json" \
    -d '{"id": 1, "name": "Jane Doe", "email": "jane.doe@example.com"}'
    ```

- **Kullanıcı silme:**

    ```sh
    curl -X DELETE http://localhost:8080/users/delete?id=1
    ```

### Testler

Testler mevcut değildir. Ancak, Go'nun test çerçevesi ile birim testleri eklemeyi düşünebilirsiniz.

### Katkıda Bulunma

Katkıda bulunmak isterseniz, lütfen bir pull request gönderin veya bir issue açın.

### Lisans

Bu proje MIT Lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasına bakabilirsiniz.

## İletişim

Proje, Uğurcan Yaş tarafından hazırlanmış ve stajyerler için geliştirilmiştir. Proje ile ilgili sorular veya geri bildirimler için [Uğurcan Yaş](mailto:uydevp@gmail.com) ile iletişime geçebilirsiniz.
