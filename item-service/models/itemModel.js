// In-memory data store for items
let items = [
  {
    id: '1',
    name: 'Laptop',
    description: 'High-performance laptop for professionals',
    price: 1299.99,
    quantity: 10,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString()
  },
  {
    id: '2',
    name: 'Wireless Mouse',
    description: 'Ergonomic wireless mouse',
    price: 29.99,
    quantity: 50,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString()
  },
  {
    id: '3',
    name: 'USB-C Cable',
    description: 'Durable USB-C charging cable',
    price: 14.99,
    quantity: 100,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString()
  }
];

let currentId = 4;

class Item {
  // Find all items
  static findAll() {
    return items;
  }

  // Find item by ID
  static findById(id) {
    return items.find(item => item.id === id);
  }

  // Create new item
  static create(itemData) {
    const newItem = {
      id: String(currentId++),
      name: itemData.name,
      description: itemData.description || '',
      price: itemData.price,
      quantity: itemData.quantity || 0,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    };
    
    items.push(newItem);
    return newItem;
  }

  // Update item
  static update(id, itemData) {
    const index = items.findIndex(item => item.id === id);
    
    if (index === -1) {
      return null;
    }

    // Update only provided fields
    const updatedItem = {
      ...items[index],
      ...itemData,
      id: items[index].id, // Keep original ID
      createdAt: items[index].createdAt, // Keep original creation date
      updatedAt: new Date().toISOString()
    };

    // Remove undefined values
    Object.keys(updatedItem).forEach(key => {
      if (updatedItem[key] === undefined) {
        updatedItem[key] = items[index][key];
      }
    });

    items[index] = updatedItem;
    return updatedItem;
  }

  // Delete item
  static delete(id) {
    const index = items.findIndex(item => item.id === id);
    
    if (index === -1) {
      return false;
    }

    items.splice(index, 1);
    return true;
  }
}

module.exports = Item;
