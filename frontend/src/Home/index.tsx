import React, {useEffect} from 'react';



const base_url_api = `${import.meta.env.VITE_REACT_APP_API_URL}`

const Home : React.FC = () => {

    

    // fetch now 

    useEffect(() => {
        document.cookie = "user=13f9cac5-1ff1-4ad3-9b88-024c4d6b9a68; path=/; expires=Thu, 01 Jan 2040 00:00:00 GMT; SameSite=None; domain=localhost"

        fetch(`${base_url_api}/application`, {
            credentials: 'include',
        })
        .then(response => response.json())
        .then(data => console.log(data))
    },[])

    return (
        <div>
            <h1>Home</h1>
        </div>
    )
}


export default Home;
