import React, {useState} from 'react';


export function App() {
    const [productName, setProductName] = useState("");
    const [productDesc, setProductDesc] = useState("");
    const [productStock, setProductStock] = useState(0);
    const [productImageURL, setProductImageURL] = useState("");

    return (
        <div className="container py-5">
            <div className="card">
                <div className="card-header">
                    <h2>Add a new product</h2>
                </div>
                 <div className="card-body">
                    <div className="row">
                        <div className="col-6">
                            <input placeholder="Product Name" type="text" className="form-control mb-1" value={productName} onChange={({target}) => setProductName(target.value)}/>
                            <input placeholder="Product Description" type="text" className="form-control mb-1" value={productDesc} onChange={({target}) => setProductDesc(target.value)}/>
                            <input placeholder="Product Stock" type="number" className="form-control mb-1" min={0} value={productStock} onChange={({target}) => setProductStock(parseInt(target.value))}/>
                           
                        </div>
                     <div className="col-6">
 <input placeholder="Product Image URL" type="text" className="form-control mb-1" value={productImageURL} onChange={({target}) => setProductImageURL(target.value)}/>
                         {productImageURL && <img src={productImageURL} className="img-fluid"/>}
                     </div>
                    </div> 


                </div>
                <div className="card-footer">
                    <button className="btn btn-primary">Submit</button>
                </div>
            </div> 
        </div>
    );
}
