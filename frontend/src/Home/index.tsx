import React, {useEffect} from 'react';
import { Link } from 'react-router-dom';
import { Application } from './types';
import CreateApplicationModel  from './create_application_modal';



const base_url_api = `${import.meta.env.VITE_REACT_APP_API_URL}`

/*
{
  "Id": "8271482f-f368-4c46-9884-0c0e27533cc6",
  "Owner": "13f9cac5-1ff1-4ad3-9b88-024c4d6b9a68",
  "Time_Created": "2024-02-10T00:50:10.783Z",
  "Time_Edited": "2024-02-10T00:50:10.783Z",
  "Name": "My First CORS Testing Application",
  "Description": "I Hope One Day I can truly understand what CORS IS DOING"
}
*/


const Home : React.FC = () => {

    const [applications, setApplications] = React.useState<Application[]>([])
    const [addModal, setAddModal] = React.useState<boolean>(false);
    const [addedApplication, setAddedApplication] = React.useState<Application>({
        Id: '',
        Owner: '',
        Time_Created: '',
        Time_Edited: '',
        Name: '',
        Description: ''
    })
    // fetch now 

   const onSubmit = async () => {

        console.log("Submitting")

        try {
           const res =await fetch(`${base_url_api}/application`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "Name": addedApplication.Name,
                    "Description": addedApplication.Description
                })
            })

            const data = await res.json()

            console.log(data)

            setApplications([...applications, data])
            setAddedApplication({
                Id: '',
                Owner: '',
                Time_Created: '',
                Time_Edited: '',
                Name: '',
                Description: ''
            })
            setAddModal(false)
            
        }
        catch (error) {
            console.log(error)
        }
    }

    const onChange = (field:string, value:string) => {
        console.log(field, value)

        if(field === 'Name') {
            setAddedApplication({
                ...addedApplication,
                Name: value
            })
        }
        if (field === 'Description') {
            setAddedApplication({
                ...addedApplication,
                Description: value
            })
        }        

    }



    useEffect(() => {
        document.cookie = "user=13f9cac5-1ff1-4ad3-9b88-024c4d6b9a68; path=/; expires=Thu, 01 Jan 2040 00:00:00 GMT; SameSite=None; domain=localhost"

        fetch(`${base_url_api}/application`, {
            credentials: 'include',
        })
        .then(response => {
            response.json()
            .then(data => {
                setApplications(data)
                console.log(data)
            })
            .catch(error => console.log(error))
        })
        .then(data => console.log(data))
    },[])



    return (
        <div className="flex w-full flex-wrap justify-center">
            <div className="w-full flex justify-center">
                <h1 className="text-2xl font-bold">Welcome to the Home Page</h1>
            </div>
            <div className="w-full flex justify-center">
                <h1 className="text-2xl font-bold">Applications</h1>
            </div>
            <div className="w-full flex justify-center">
                <button onClick={()=> setAddModal(true)} className="bg-blue-500 text-white p-2 rounded-md">Add Application</button>
            </div>

           
                
            <div className="w-full flex justify-center flex-wrap">
                <h1 className="text-2xl font-bold">Applications</h1>
            
            {
                addModal?
                CreateApplicationModel({
                    cancel: ()=> setAddModal(false),
                    onChange: onChange,
                    state: addedApplication,
                    submission: onSubmit
                })
                :<div className='w-full flex justify-center'>
                {
                    applications.map((application, index) => {
                        return (
                            <div key={index} className="w-1/3 flex justify-center">
                                <div className="w-1/2 flex justify-center flex-wrap">
                                    <h1 className="text-xl font-bold">{application.Name}</h1>
                                    <p className="text-lg">{application.Description}</p>
                                    <Link to={`/application/${application.Id}`} className="bg-blue-500 text-white p-2 rounded-md inline-block">View</Link>
                                    <button className="bg-blue-500 text-white p-2 rounded-md">View</button>
                                </div>
                            </div>
                        )
                    })
                }
            </div>
            }

        </div>

        </div>
    )
}


export default Home;
