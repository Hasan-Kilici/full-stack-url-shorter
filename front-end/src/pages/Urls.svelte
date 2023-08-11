<script>
    let Links = [];

    function GetLinks(){
        fetch("http://localhost:3000/list/user/urls",{
            credentials: 'include',
        }).then(async(res)=>{
            let data = await res.json();
            Links = data.data;
        })
    }

    GetLinks()
    
</script>
{#each Links as Link}
    <a href="/clicks/{Link.Token}">{Link.Redirect}</a><br>
{/each}