import React, {useCallback, useState} from 'react';


export function App() {
    const [productName, setProductName] = useState("");
    const [productDesc, setProductDesc] = useState("");
    const [productStock, setProductStock] = useState(0);
    const [productImageURL, setProductImageURL] = useState("");
    const deps = [productName, productDesc, productStock, productImageURL];

    const reset = useCallback(() => {
        setProductName("");
        setProductDesc("");
        setProductStock(0);
        setProductImageURL("");
    }, []);

    const handleSubmit = useCallback(() => {
        if(deps.some(f => !f)) {
            alert("All fields are required!");
            return;
        }

        fetch("https://product.free.beeceptor.com/product", {
            method: "POST",
            body: JSON.stringify({
                productName,
                productDesc,
                productStock,
                productImageURL
            }),
            headers: {"Content-Type": "application/json"}
        }).then(reset).catch(error => {
            reset();
            alert(error)
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
                            <input placeholder="Product Name" type="text" className="form-control mb-1" value={productName} onChange={({target}) => setProductName(target.value)}/>
                            <input placeholder="Product Description" type="text" className="form-control mb-1" value={productDesc} onChange={({target}) => setProductDesc(target.value)}/>
                            <input placeholder="Product Stock" type="number" className="form-control mb-1" min={0} value={productStock} onChange={({target}) => setProductStock(parseInt(target.value))}/>
                           
                        </div>
                     <div className="col-6">
                         <input 
                            placeholder="Product Image URL" 
                            type="text" 
                            className="form-control mb-1" 
                            value={productImageURL} 
                            onChange={({target}) => setProductImageURL(target.value)}/>
                         {productImageURL && <img src={productImageURL} className="img-fluid"/>}
                     </div>
                    </div> 


                </div>
                <div className="card-footer text-right">
                    <button className="btn btn-primary" onClick={handleSubmit}>Submit</button>
                </div>
            </div> 
        </div>
    );
}
