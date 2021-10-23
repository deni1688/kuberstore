import React, {ChangeEvent, useCallback, useState} from 'react';


const dataDefault = {
    name: "",
    desc: "",
    stock: 0,
    imageURL: ""
};

export function App() {
    const [data, setData] = useState(dataDefault);
    const [msg, setMessage] = useState({type: "", text: ""});
    const deps = Object.values(data);

    const setInput = ({target}:ChangeEvent<HTMLInputElement>) => setData({...data, [target?.name]: target?.value});
    const showMessage = (text: string, type: string = "info") => {
        setMessage({text, type});
        setTimeout(() => setMessage({text: "", type: ""}), 3000);
    };

    const handleSubmit = useCallback(() => {
        if(deps.some(f => !f)) {
            showMessage("All fields are required!");
            return;
        }

        fetch("http://kubestore.de/products", {
            method: "POST",
            body: JSON.stringify(data),
            headers: {"Content-Type": "application/json"}
        }).then(res => res.json()).then(res => {
            setData(dataDefault)
            showMessage(res.message, "success")
        }).catch(err => {
            setData(dataDefault)
            showMessage(err.message, "danger");
        });
    },deps);

    return (
        <div className="container py-5">
            <div className="card">
                <div className="card-header">
                    <h2>Product Admin</h2>
                </div>
                 <div className="card-body">
                    <div className="row">
                        <div className="col-6">
                            <input
                                placeholder="Product Name"
                                type="text"
                                name="name"
                                className="form-control mb-1"
                                value={data.name}
                                onChange={setInput}/>
                            <input
                                placeholder="Product Description"
                                type="text"
                                name="desc"
                                className="form-control mb-1"
                                value={data.desc} onChange={setInput}/>
                            <input
                                placeholder="Product Stock"
                                type="number"
                                name="stock"
                                className="form-control mb-1"
                                min={0}
                                value={data.stock}
                                onChange={({target}) => setData({...data, [target.name]: parseInt(target.value)})}/>
                        </div>
                     <div className="col-6">
                         <input
                            placeholder="Product Image URL"
                            type="text"
                            className="form-control mb-1"
                            name="imageURL"
                            value={data.imageURL}
                            onChange={setInput}/>
                         {data.imageURL && <img src={data.imageURL} className="img-fluid"/>}
                     </div>
                    </div>


                </div>
                <div className="card-footer d-flex justify-content-between align-items-center">
                    <span className={"text-"+msg.type}>{msg.text}</span>
                    <button className="btn btn-primary" onClick={handleSubmit}>Submit</button>
                </div>
            </div>
        </div>
    );
}
