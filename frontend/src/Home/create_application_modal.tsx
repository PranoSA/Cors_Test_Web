import { Application } from "./types";

type CreateApplicationModelProps = {
    submission: () => void;
    onChange : (field:string, value:string) => void;
    state : Application;
    cancel : () => void;
}


const CreateApplicationModel= (props:CreateApplicationModelProps) => {



    return (
        <div className='w-full flex flex-wrap'>
            <div className="w-full flex justify-center p-5">
                <h1 className='text-center text-3xl '>Create Application</h1>
            </div>
            
            <div className="w-1/2 flex justify-center p-5">
                <button onClick={()=> props.cancel()} className="bg-red-500 text-2xl text-white p-2 py-5 px-5 rounded-md">Cancel</button>
            </div>

            <div className="w-1/2 flex justify-center p-5">
                <button onClick={()=> props.cancel()} className="bg-blue-500 text-2xl text-white p-2 py-5 px-5 rounded-md"> Submit </button>
            </div>

            <div className="w-full flex justify-center p-5">
                <input onChange={(e)=> props.onChange('Name', e.target.value)} value={props.state.Name} className="w-1/2 p-2 text-2xl border-2 border-blue-500 rounded-md" type="text" placeholder="Application Name" />
            </div>

            <div className="w-full flex justify-center p-5">
                <input onChange={(e)=> props.onChange('Description', e.target.value)} value={props.state.Description} className="w-1/2 p-2 text-2xl border-2 border-blue-500 rounded-md" type="text" placeholder="Application Description" />
            </div>

        </div>
    )
}

export default CreateApplicationModel
