<script>
  import { Router, Link, Route } from "svelte-routing";
  import Home from "./pages/Home.svelte";
  import Register from "./pages/Register.svelte";
  import Login from "./pages/Login.svelte";
  import CreateUrl from "./pages/CreateLink.svelte"
  import Redirect from "./pages/Redirect.svelte";
  import Urls from "./pages/Urls.svelte";
  import Clicks from "./pages/Clicks.svelte";
  
  export let url = "/";
  import Cookies from "js-cookie";
  let User = false;
  let user = Cookies.get("Token");
  async function GetUser(){
  url = window.location.pathname;
  if(user){
    const response = await fetch(`http://localhost:3000/find/user/${user}`, {
        headers: {
          "Content-Type": "application/json",
        }
      });
  
      if (response.ok) {
        const data = await response.json();
        User = data.data;
        console.log(data)
      } else if (response.status === 404) {
        console.log("Kullanıcı bulunamadı.");
      }
    } 
  }
  GetUser();
  
  </script>
  
  
  <Router {url}>
  <nav>
    <Link to="/">Anasayfa</Link>
    {#if !User}
    <Link to="/register">Kayıt ol</Link>
    <Link to="/login">Giriş yap</Link>
    {:else}
    {User.Username}
      <Link to="/create/url">Link kısalt</Link>
      <Link to="/my/urls">Linklerim</Link>
      <Route path="/create/url" component={CreateUrl} />
      <Route path="/clicks/:UrlToken" let:params>
        <Clicks UrlToken={params.UrlToken}/>
    </Route>
    {/if}
  </nav>
  <Route path="/" component={Home} />
  <Route path="/register" component={Register} />
  <Route path="/login" component={Login} />
  <Route path="/r/:token" let:params>
      <Redirect token={params.token}/>
  </Route>
  <Route path="/clicks/:UrlToken" let:params>
    <Clicks UrlToken={params.UrlToken}/>
</Route>
  <Route path="/my/urls" component={Urls} />
  </Router>