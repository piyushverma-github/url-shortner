# 🚀 GO URL Shortener

A fast and simple **URL shortener** built with **Go** for the backend and **HTML/CSS/JavaScript** for the frontend.  
Shorten your long, messy URLs into clean, shareable links — in just one click.

🌍 **Live Demo:** (https://url-shortner-gmda.onrender.com)  

---

## ✨ Features
- ⚡ **Fast** — Shortens URLs instantly.
- 🔒 **Safe** — Your original URLs are kept private.
- 🖥 **Frontend + Backend** in one — No complicated setup.
- 📱 **Responsive UI** — Works on desktop & mobile.
- 🔄 **Direct Redirects** — Use short links like `/abc123` to jump to your destination.

---

## 📸 Screenshots

### Homepage
![Homepage]<img width="1878" height="808" alt="image" src="https://github.com/user-attachments/assets/ae86b51d-5148-471d-8738-ff760560398f" />


### Shortened Link Example
![Shortened URL]<img width="1876" height="483" alt="image" src="https://github.com/user-attachments/assets/2f037dec-e2cd-4733-89d3-f4f16ff997aa" />


---

## 📦 Tech Stack
- **Backend:** Go (net/http, crypto/md5)
- **Frontend:** HTML, CSS, JavaScript (Fetch API)
- **Hosting:** Render

---

## 🛠 Installation & Setup

### 1️⃣ Clone this repo
```bash
git clone https://github.com/piyushverma-github/url-shortner.git
cd url-shortner

2️⃣ Install dependencies
Go doesn’t need package managers for standard libs used here.

3️⃣ Run the server
```bash
go run main.go

4️⃣ Open in browser
```arduino
http://localhost:3000

🌐 How It Works
1. Enter your long URL in the input field.
2. Click "Shorten" — The backend generates a 6-character hash using MD5.
3. Get your short link — Click to visit instantly.
4. Share it — Your friends can use it directly.

📂 Project Structure
url-shortner/
├── go.mod              # Go module file
├── main.go             # Backend server
├── public/             # Frontend assets
│   ├── index.html      # Homepage
│   ├── script.js       # Frontend logic
│   └── style.css       # Styling
└── package.json        # (Optional) Node.js metadata

🤝 Contributing
Pull requests are welcome! If you find bugs or have feature requests,
feel free to open an issue or fork and improve.

👨‍💻 Author
Made with ❤️ by Piyush Verma
You can also go throught the other projects
