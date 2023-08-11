<script>
    export let token

    let Header;

    async function Redirect(){
        fetch("https://freeipapi.com/api/json").then(async (data) => {
            let res = await data.json();
            Header = {
                ip : res.ipAddress,
                country : res.countryCode,
                city : res.cityName,
                latitude : String(res.latitude),
                longitude : String(res.longitude),
            }
            console.log(Header)
        }).then(async(send)=>{
            CallApi()
        })
    }

    function CallApi(){
        fetch(`http://localhost:3000/r/${token}`, {
            method:"POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(Header),
        }).then(async(res)=>{
            let response = await res.json();
            console.log(response)
            setTimeout(async()=>{
            window.location.href = response.redirect;
            },1500)
        })
    }
    Redirect()
</script>
<h3>Bekle</h3>