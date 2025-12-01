# RetroTVIndia Backend ⚙️📺🇮🇳

The **RetroTVIndia Backend** powers the core functionality of the RetroTVIndia platform — a nostalgic recreation of classic Indian television inspired by myretrostv.com.  

This backend provides structured APIs for serving decade-based playlists, categories, metadata, and YouTube video links to the frontend. Lightweight, fast, and built for scalability.

## 🚀 Features
- 📂 Serves curated JSON data for shows, categories, decades, and YouTube URLs  
- ▶️ Handles playlist and episode metadata for classic Indian TV content  
- ⚡ Fast API responses optimized for static + dynamic content  
- 🌐 CORS-ready and frontend friendly  
- 🧩 Simple, modular, and easy to extend  

## 🛠️ Tech Stack
- **Language:** Go (Golang) 🐹  
- **Framework:** net/http / chi / gin (based on project setup)  
- **Data Storage:** JSON files  
- **Deployment:** Lightweight server-friendly (Raspberry Pi compatible)

## ▶️ Running Locally

Follow these steps to get the backend server running on your machine:
```bash
git clone https://github.com/RetroTVIndia/backend.git
cd backend
go mod download
go run main.go
```
> [!NOTE]
> Your backend will start at: *http://localhost:8080*
> (Ensure this port is free before running)

## 📡 API Endpoints
<details>
<summary><code>GET</code> <b>/</b></summary>
<br>
Health check to verify the server is running.
<br><br>
<b>Response:</b>
<pre>
{
  "message": "Server is running!"
}
</pre>
</details>

<details>
<summary><code>GET</code> <b>/categories</b></summary>
<br>
Returns a list of all available content categories.
<br><br>
<b>Response:</b>
<pre>
[
  "Music",
  "Commercials",
  "Trailers"
]
</pre>
</details>

<details>
<summary><code>GET</code> <b>/category</b></summary>
<br>
Returns a list of all videos belonging to a specific category.
<br><br>
<b>Query Parameters:</b>
<ul>
    <li><code>name</code> (Required): The exact name of the category.</li>
</ul>
<b>Example:</b> <code>/category?name=Commercials</code>
</details>

<details>
<summary><code>GET</code> <b>/random</b></summary>
<br>
Returns a single random video.
<br><br>
<b>Query Parameters:</b>
<ul>
    <li><code>category</code> (Optional): Filter the random selection by specific categories. You can repeat this parameter to include multiple categories.</li>
</ul>
<b>Example:</b> <code>/random?category=Music&category=Trailers</code>
<br>
<i>If no category is provided, it selects from the entire database.</i>
</details>

## 🤝 Contributing
We welcome contributions! Whether you want to add new shows, fix metadata, or improve the UI/UX, your help will make RetroTVIndia even better.  

1. Fork the repository 🍴  
2. Create your feature branch (`git checkout -b feature-name`) 🌿  
3. Commit your changes (`git commit -m 'Add some feature'`) 💾  
4. Push to the branch (`git push origin feature-name`) 🚀  
5. Open a Pull Request 🔃  

## 📬 Contact
For inquiries or collaboration opportunities, reach out via GitHub Issues or email at **aglakshya06@gmail.com**.  

## ⚖️ License
This project is licensed under the **MIT License**.  

🎉 RetroTVIndia Backend — powering nostalgia, one episode at a time!
