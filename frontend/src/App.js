import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
    const [items, setItems] = useState([]);
    const [selectedItem, setSelectedItem] = useState(null);
    const [newName, setNewName] = useState('');
    const [newPrice, setNewPrice] = useState('');

    useEffect(() => {
        fetchItems();
    }, []);

    const fetchItems = () => {
        axios.get('http://localhost:8080/items')
            .then(response => setItems(response.data))
            .catch(error => console.error('Error loading items:', error));
    };

    const handleUpdateClick = (item) => {
        setSelectedItem(item);
        setNewName(item.name);
        setNewPrice(item.price);
    };

    const handleUpdateItem = () => {
        const updatedItem = {
            ...selectedItem,
            name: newName,
            price: parseFloat(newPrice)
        };

        axios.put(`http://localhost:8080/item?id=${selectedItem.id}`, updatedItem)
            .then(() => {
                fetchItems();
                setSelectedItem(null);
                setNewName('');
                setNewPrice('');
            })
            .catch(error => console.error('Error updating item:', error));
    };

    const handleDeleteItem = (id) => {
        axios.delete(`http://localhost:8080/item?id=${id}`)
            .then(() => fetchItems())
            .catch(error => console.error('Error deleting item:', error));
    };

    return (
        <div>
            <h1>Items</h1>
            <ul>
                {items.map(item => (
                    <li key={item.id}>
                        {item.name} - ${item.price}
                        <button onClick={() => handleUpdateClick(item)}>Update</button>
                        <button onClick={() => handleDeleteItem(item.id)}>Delete</button>
                    </li>
                ))}
            </ul>

            {selectedItem && (
                <div>
                    <h2>Update Item</h2>
                    <label>
                        Name:
                        <input 
                            type="text" 
                            value={newName} 
                            onChange={e => setNewName(e.target.value)} 
                        />
                    </label>
                    <label>
                        Price:
                        <input 
                            type="number" 
                            value={newPrice} 
                            onChange={e => setNewPrice(e.target.value)} 
                        />
                    </label>
                    <button onClick={handleUpdateItem}>Save</button>
                </div>
            )}
        </div>
    );
}

export default App;
