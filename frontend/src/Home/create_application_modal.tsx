import { Application } from "./types";

type CreateApplicationModelProps = {
    submission: () => void;
    onChange : (field:string, value:string) => void;
    state : Application;
    cancel : () => void;
}


const CreateApplicationModel= (props:CreateApplicationModelProps) => {



    return (
        <div>
            <h1>Create Application</h1>
            <button onClick={()=> props.cancel()}>Cancel</button>
            <input type="text" value={props.state.Name} onChange={(e)=> props.onChange('Name', e.target.value)} placeholder="Name"/>
            <input type="text" onChange={(e)=> props.onChange('Description', e.target.value)} placeholder="Description"/>
            <button onClick={props.submission}> SUBMIT APPICATION </button>
        </div>
    )
}

export default CreateApplicationModel
