<script>
    import Cookies from "js-cookie";
  
    let email = "";
    let pass = "";
    let error = "";
  
    async function login() {
      const formData = {
        email: email,
        pass: pass,
      };
  
      try {
        const response = await fetch("http://localhost:3000/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(formData),
        });
  
        if (response.ok) {
          let lastPath = Cookies.get("LastPath")
          const data = await response.json();
          const token = data.Token;
  
          Cookies.set("Token", token);
          console.log("Giriş başarılı! Token:", token);
          window.location.href = lastPath
  
        } else if (response.status === 404) {
          error = "Kullanıcı bulunamadı.";
        } else if (response.status === 400) {
          error = "E-posta ya da şifre yanlış.";
        } else {
          error = "Bir hata oluştu.";
          console.error("Bir hata oluştu.");
        }
      } catch (error) {
        console.error("İstek gönderilirken hata oluştu:", error);
      }
    }

  </script>
  
  <main>
    <h1>Giriş Yap</h1>
    {#if error}
      <p style="color: red;" class="text-red-400 font-extrabold text-center">{error}</p>
    {/if}
   <div class="p-4 mb-1.5">
     <form class="roundex-3xl bg-gray-300 py-5 px-5">
      <label for="email" class="text-blue-600 pt-4">E-posta:</label>
      <input type="email" bind:value={email} id="email">
  
      <label for="pass" class="text-blue-600 pt-4">Şifre:</label>
      <input type="password" bind:value={pass} id="pass">
  
      <button on:click|preventDefault={login}>Giriş Yap</button>
    </form>
   </div>
  </main>