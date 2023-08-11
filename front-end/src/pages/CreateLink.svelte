<script>
    let redirect = "";
    let message;
    async function createUrl() {
        const response = await fetch("http://localhost:3000/create/url", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
            body: JSON.stringify({
                redirect: redirect,
            }),
        }).then((data)=>{
            if(data.ok){
                message = "Link başarıyla kısaltıldı"
            }
        })
    }
</script>
<main>
    {#if message}
        {message}
    {/if}
    <input type="text" bind:value={redirect}>
    <button on:click={createUrl}>Oluştur</button>
</main>