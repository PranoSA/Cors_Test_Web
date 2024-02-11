import React, {useEffect} from 'react';
import { Link , useNavigate} from 'react-router-dom';
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
    const [editModal, setEditModal] = React.useState<boolean>(false);
    const [addedApplication, setAddedApplication] = React.useState<Application>({
        Id: '',
        Owner: '',
        Time_Created: '',
        Time_Edited: '',
        Name: '',
        Description: ''
    })
    // fetch now 
    const  navigate = useNavigate();

   const onSubmit = async () => {

        console.log("Submitting")

        try {
           const res =await fetch(`${base_url_api}/application`, {
                method: editModal? 'PUT' : 'POST',
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

    const openEditModal = (index:number) => {
        setAddedApplication(applications[index])
        setEditModal(true)
    }


    return (
        <div className="flex w-full flex-wrap justify-center">
            <div className="w-full flex justify-center p-5">
                <h1 className=" text-3xl font-bold"> Application View : </h1>
            </div>
            <div className="w-full flex justify-center p-5">
                <button onClick={()=> setAddModal(true)} className="bg-blue-500 text-2xl text-white p-2 py-5 px-5 rounded-md">Add Application</button>
            </div>

           
            
            {
                (addModal||editModal)?
                CreateApplicationModel({
                    cancel: ()=> {
                        setAddModal(false)
                        setEditModal(false)
                    },
                    onChange: onChange,
                    state: addedApplication,
                    submission: onSubmit
                })
                :<div className='w-full flex flex-wrap justify-center'>
                {
                    applications.map((application, index) => {
                        return (
                            <div key={index} className="w-1/3 flex justify-center border border-5 border-red-500 p-20">
                                <div className="w-full flex justify-center flex-wrap">
                                    <div className='w-full justify-center'>
                                        <h1 className="text-3xl font-bold text-center">{application.Name}</h1>    
                                    </div>
                                    <div className='w-full text-center'>
                                        <h1 className='w-full text-center text-2xl font-bold pt-5'> Description : </h1>
                                    </div>
                                    <div className='w-full text-center text-lg p-3'>
                                        <p className="text-lg">{application.Description}</p>
                                    </div> 
                                    <div className='w-full text-center text-lg p-3'>
                                        <p className='text-2xl font-bold pr-3'> Created At : </p>
                                        <p className="text-lg"> {application.Time_Created.split('.')[0].split("T").join("  ")}</p>
                                    </div>

                                    <div className='w-full text-center text-lg p-3'>
                                        <p className='text-2xl font-bold pr-3'> Lasted Modified : </p>
                                        <p className="text-lg"> {application.Time_Edited.split('.')[0].split("T").join("  ")}</p>
                                    </div>
                                    <div className='w-full flex justify-around items-center'>
                                        <Link to={`/application/${application.Id}`} className="bg-blue-500 text-white p-2 rounded-md w-32 h-12">View</Link>
                                        <button className="bg-blue-500 text-white p-2 rounded-md w-32 h-12" onClick={() => navigate(`/application/${application.Id}`)}>View</button>
                                        <button className="bg-blue-500 text-white w-32 h-12 rounded-md inline-block" onClick={() => openEditModal(index)}> Edit</button>                                        
                                    </div >
                                </div>
                            </div>
                        )
                    })
                }
            </div>
            }


        </div>
    )
}


export default Home;
