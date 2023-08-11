<script>
    import Cookies from "js-cookie";
    let error = "";
  
    async function ValidateEmail(mail) {
      return /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(mail);
    }
  
    function ValidatePassword(password) {
      if (password.includes(" ")) {
        error = "Şifre boşluk içeremez";
      } else if (password.length < 8) {
        error = "Şifre en az 8 karakter olmak zorundadır.";
      } else if (password.length > 32) {
        error = "Şifre en fazla 32 karakter olmalıdır.";
      } else if (!/\d/.test(password)) {
        error = "Şifrede sayılar yer almalıdır.";
      } else if (!/[a-zA-Z]/.test(password)) {
        error = "Şifrede metin bulunmak zorundadır.";
      } else if (!/[!@#$%^&*(),.?":{}\|_<>]/.test(password)) {
        error = "Şifrede özel karakter bulunmak zorundadır.";
      } else {
        error = "";
      }
    }
  
    function ValidateUsername(username) {
      return !username.includes(" ");
    }
  
    async function Register() {
      const username = document.getElementById("username").value;
      const email = document.getElementById("email").value;
      const password = document.getElementById("pass").value;
  
      if (!username) {
        error = "Kullanıcı adı boş bırakılamaz.";
      } else if (!ValidateUsername(username)) {
        error = "Kullanıcı adı boşluk içermemelidir.";
      } else if (!ValidateEmail(email)) {
        error = "Bu Email geçersiz.";
      } else {
        ValidatePassword(password);
      }
  
      if (!error) {
        const formData = {
          username: username,
          email: email,
          pass: password,
        };
  
        try {
          const response = await fetch("http://localhost:3000/register", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(formData),
          });
  
          if (response.ok) {
            let lastPath = Cookies.get("LastPath")
            console.log("Kullanıcı başarıyla kaydedildi");
            window.location.href = lastPath;
            let res = await response.json()
            Cookies.set("Token", res.Token);
          } else if (response.status === 401) {
            error = "Yetkilendirme hatası: Kullanıcı girişi gerekiyor.";
          } else if (response.status === 403) {
            console.error("Erişim reddedildi: Yönetici izni gerekiyor.");
          } else {
            console.error("Bir hata oluştu.");
            error = "Bu Email Kullanılıyor.";
          }
        } catch (error) {
          console.error("İstek gönderilirken hata oluştu:", error);
        }
      }
    }
  </script>
  <main>
    {#if error}
      <div id="error" class="text-red-400 font-extrabold text-center">{error}</div>
    {/if}
  
    <div>
     <input type="text" id="username" placeholder="Kullanıcı adı" />
     <div></div>
     <input type="text" id="email" placeholder="Email"/>
     <div></div>
     <input type="password" id="pass" placeholder="Şifre"/>
      <div class="pt-4"></div>
    <button on:click={Register}>Kaydol</button>
   </div>
  </main>
  
  <style>
    #error {
      display: block;
      color: red;
      margin-bottom: 10px;
    }
  
    input,
    button {
      margin-bottom: 10px;
    }
  </style>