<script>
    export let UrlToken;
    let Clicks = [];
    async function ListClicks(){
        fetch(`http://localhost:3000/list/clicked/${UrlToken}`).then(async(data)=>{
            let res = await data.json();
            Clicks = res.data;
        })
    }
    function deleteUrl(){
        fetch(`http://localhost:3000/delete/url/${UrlToken}`,{
            method:"DELETE",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
        }).then(async(data)=>{
            if(data.ok){
                window.location.href = "/"
            }
        })
    }
    ListClicks()
</script>
<button on:click={deleteUrl}>Bu linki sil</button>
<p>Link : http://localhost:5173/r/{UrlToken}</p>
<table>
    <h3>{Clicks.length} Tıklanma</h3>
    <tr>
        <th>Ülke</th>
        <th>Şehir</th>
        <th>IP Adresi</th>
    </tr>
    {#each Clicks as Click}
        <tr>
            <td><img style="border-radius:8px;" src="https://www.countryflagicons.com/FLAT/64/{Click.Country}.png"></td>
            <td>{Click.City}</td>
            <td>{Click.IP}</td>
        </tr>
    {/each}
</table>