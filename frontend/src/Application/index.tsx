import React, { useEffect } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import { Application, CorsTests } from '../Home/types';
import AddTestPanel from './AddTestPanel';


const base_url_api = `${import.meta.env.VITE_REACT_APP_API_URL}`

const ApplicationDisplay: React.FC = () => {

    
    //Get ID from URL
    const url = window.location.href;
    const {id} = useParams<{id:string}>();

    const [application, setApplication] = React.useState<Application|null>(null)
    const [tests, setTests] = React.useState<CorsTests[]>([])
    const [addModal, setAddModal] = React.useState<boolean>(false);

    const [loadingTestResults, setLoadingTestResults] = React.useState<boolean>(false)
    const [addedTest, setAddedTest] = React.useState<CorsTests>({
        Id: '',
        Owner: '',
        Origin: '',
        ApplicationId: '',
        Endpoint: '',
        Methods: '',
        Headers: '',
        Authentication: false,
        Created_at: ''
    })

    const navigator = useNavigate();

    const handleSubmission = async () => {

        console.log("Submitting")

        try {
            const res =await fetch(`${base_url_api}/test/${id}`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "ApplicationId": id,
                    "Origin": addedTest.Origin,
                    "Endpoint": addedTest.Endpoint,
                    "Methods": addedTest.Methods,
                    "Headers": addedTest.Headers,
                    "Authentication": addedTest.Authentication,
                })
            })

            const data = await res.json()

            console.log(data)

            setTests([...tests, data])
            setAddedTest({
                Id: '',
                Owner: '',
                Origin: '',
                ApplicationId: '',
                Endpoint: '',
                Methods: '',
                Headers: '',
                Authentication: false,
                Created_at: ''
            })
            setAddModal(false)
        }
        catch (error) {
            console.log(error)
        }  
    }

    const handleCancel = () => {
        setAddModal(false)
    }

    const handleChange = (field:string, value:string) => {
        console.log(field, value)

        if(field === 'Origin') {
            setAddedTest({
                ...addedTest,
                Origin: value
            })
        }
        if (field === 'Endpoint') {
            setAddedTest({
                ...addedTest,
                Endpoint: value
            })
        }
        if (field === 'Methods') {
            setAddedTest({
                ...addedTest,
                Methods: value
            })
        }
        if (field === 'Headers') {
            setAddedTest({
                ...addedTest,
                Headers: value
            })
        }
        if (field === 'Authentication') {
            setAddedTest({
                ...addedTest,
                Authentication: value === 'true' ? true : false
            })
        }
    }


    useEffect(() => {
        console.log(id)

        const fetchApplication = async () => {
            try {
                const res = await fetch(`${base_url_api}/application/${id}`, {
                    method: 'GET',
                    credentials: 'include'
                })

                const data = await res.json()

                console.log(data)

                setApplication(data)
            }
            catch (error) {
                console.log(error)
            }
        }

        fetchApplication()
    },[id])

    useEffect(() => {
        const fetchTests = async () => {
            try {
                const res = await fetch(`${base_url_api}/tests/${id}`, {
                    method: 'GET',
                    credentials: 'include'
                })

                const data = await res.json()

                console.log(data)

                setTests(data)
            }
            catch (error) {
                console.log(error)
            }
        }

        fetchTests()
    },[id])

    if(application === null) {
        return (
            <div>
                <h1>Loading...</h1>
            </div>
        )
    }

    const TestsPanel = () => {


        if (tests.length === 0) {
            return (
                <div>
                    <h1>No Tests</h1>
                </div>
            )
        }
        return (
            <div className='w-full flex flex-wrap'>
                    {tests.map((test, index) => {
                        return (
                            <div key={index}>
                                <h1>{test.ApplicationId}</h1>
                                <p>{test.Authentication}</p>
                                <p>{test.Created_at}</p>
                                <p>{test.Endpoint}</p>
                                <p>{test.Id}</p>
                                <p>{test.Methods}</p>
                                <p>{test.Origin}</p>
                                <p> {test.Headers}</p>
                            </div>
                        )
                    })
                }
            </div>
        );
}

    if(addModal) {
        return (
            <div className='w-full flex flex-wrap'>
                {
                AddTestPanel({
                    submission: handleSubmission,
                    onChange: handleChange,
                    state: addedTest,
                    cancel: handleCancel
                })
                }
            </div>
        )
    }


    const runTests = async() => {

        ///results/8271482f-f368-4c46-9884-0c0e27533cc6
        ///result/8271482f-f368-4c46-9884-0c0e27533cc6

        setLoadingTestResults(true)

        try {
            const res = await fetch(`${base_url_api}/result/${id}`, {
                method: 'POST',
                credentials: 'include'
            })

            const data = await res.json()

            console.log(data)

            setLoadingTestResults(false)
        }
        catch (error) {
            console.log(error)
        }

    }


    return (
        <div className='w-full flex flex-wrap justify-around'>
            <div className='w-full flex justify-around p-30'>
                
                <button className='bg-blue-500 w-40 h-12 rounded-md' onClick={()=> window.location.href = '/'}>Back</button>
                <button className='bg-blue-500 w-40 h-12 rounded-md' onClick={()=> setAddModal(true)}>Add Test</button>
                <button className='bg-blue-500 w-40 h-12 rounded-md' onClick={()=> navigator(`/results/${application.Id}`)}>View Results</button>
                <button className='bg-blue-500 w-40 h-12 rounded-md' onClick={runTests}>Run Tests</button>
            </div> 

            <div className='w-full flex justify-center p-5'>
                <h1 className='text-3xl text-center'> Application Details </h1>
            </div>
            <div className='w-full flex justify-center p-5'>
                <div className='w-1/2 flex flex-wrap justify-center p-5'>
                    <div className='w-full flex justify-center'>
                        <h1 className='text-3xl'>Application Name : </h1>
                    </div>
                    <div className='w-full flex justify-center'>
                        <p className='text-2xl'>{application?.Name}</p>
                    </div>

                    <div className='w-full flex justify-center p-5'>
                        <h1 className='text-3xl'>Application Description : </h1>
                        </div>
                    <div className='w-full flex justify-center'>
                        <p className='text-2xl'>{application?.Description}</p>
                    </div>

                    <div className='w-full flex justify-center p-5'>
                        <h1 className='text-3xl'>Creation Date : </h1>
                    </div>

                    <div className='w-full flex justify-center'>
                        <p className='text-2xl'>{application?.Time_Created.split(".")[0].split("T").join("  ")}</p>
                    </div>

                    <div className='w-full flex justify-center p-5'>
                        <h1 className='text-3xl'>Last Updated : </h1>
                    </div>

                    <div className='w-full flex justify-center'>
                        <p className='text-2xl'>{application?.Time_Edited.split(".")[0].split("T").join("  ")}</p>
                    </div>

                </div>
            </div>

            <div className='w-full flex justify-center p-5'>
                <div className='w-full flex justify-center text-center'>
                    <h1 className='text-3xl'>Tests</h1>
                </div>

                <div className='w-full flex justify-center p-5'>
                    {TestsPanel()}
                </div>
            </div>
        </div>
    )
}

export default ApplicationDisplay;