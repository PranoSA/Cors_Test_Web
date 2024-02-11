import React from 'react';

import {CorsTestResults} from '../Home/types'
import { useParams } from 'react-router-dom';

const ViewResults: React.FC = () => {

    const [expanded, setExpanded] = React.useState<boolean[]>([]);

    const [results, setResults] = React.useState<CorsTestResults[]>([]);

    const {id} = useParams<{id:string}>();

    React.useEffect(() => {
        const fetchResults = async () => {
            try {
                const res = await fetch(`${import.meta.env.VITE_REACT_APP_API_URL}/results/${id}`,{
                    credentials: 'include'
                })
                const data = await res.json()
                setResults(data)
            }
            catch (error) {
                console.log(error)
            }
        }
        fetchResults()
    }, [])

    if (results.length === 0) {
        return (
            <div>
                <h1>No Results Found / Loading </h1>
            </div>
        )
    }

    return (
        <div>
            {
                results.map((result, index) => {
                    return (
                        <div key={result.Id} className={`p-4 ${result.Okay ? 'bg-green-200' : 'bg-red-200'}`}>
                    
                            <h1 onClick={()=> setExpanded([...expanded, !expanded[index]])}>{result.Origin}:{result.Header}:{result.Method}</h1>
                            {
                                expanded[index] && (
                                    <div>
                                        <p>Endpoint: {result.Endpoint}</p>
                                        <p>Methods: {result.Method}</p>
                                        <p>Headers: {result.Header}</p>
                                        <p>Authentication: {result.Authentication}</p>
                                        <p>Okay: {result.Okay}</p>
                                        <p>Errors: {result.Errors}</p>
                                        <p>Access-Control-Allow-Origin: {result.Return_Access_Control_Allow_Origin}</p>
                                        <p>Access-Control-Allow-Methods: {result.Return_Access_Control_Allow_Method}</p>

                                        <p> </p>
                                    </div>
                                )
                            }
                        </div>
                    )
                })
            }
        </div>
    )
}

export default ViewResults;