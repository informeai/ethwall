
async function run(){
    try{
        let content = document.getElementById('content')
        let card = document.getElementById('card')
        let address = document.getElementById('address')
        let snack = document.getElementById('snack')
        let footer = document.getElementById('footer')
        let titleHead = document.getElementById('title-head')
        let data = {
            pass_phrase: 'secret'
        }
        fetch('http://localhost:4000/wallet',{
            method: 'POST',
            cache: 'no-cache',
            headers: {
                'Content-Type': 'application/json'
            },
            body:JSON.stringify(data)
        })
        .then(response => response.json())
        .then(data => {
            address.innerHTML = data.address
            content.classList.remove('inactive')
            card.classList.remove('inactive')
            address.classList.remove('inactive')
            footer.classList.remove('inactive')
            titleHead.classList.remove('inactive')
        })
        .catch(e => console.log(e))
        
        address.addEventListener('click',(event)=>{
            address.focus()
            navigator.clipboard.writeText(address.innerText)
            .then(()=>{
                snack.classList.remove('inactive')
                setTimeout(()=> snack.classList.add('inactive'),2000)
            })
            .catch(e => console.log(e))
        })

        card.addEventListener('focus',(event)=>{
            fetch('http://localhost:4000/wallet',{
            method: 'POST',
            cache: 'no-cache',
            headers: {
                'Content-Type': 'application/json'
            },
            body:JSON.stringify(data)
            })
            .then(response => response.json())
            .then(data => {
                address.innerHTML = data.address
                card.classList.remove('inactive')
                address.classList.remove('inactive')
            })
            .catch(e => console.log(e))
        })

    }catch(e){
        console.log(e)
    }

}

run()
