const Item = require('../models/itemModel');

// Get all items
const getAllItems = (req, res) => {
  try {
    const items = Item.findAll();
    res.status(200).json({
      success: true,
      count: items.length,
      data: items
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
};

// Get item by ID
const getItemById = (req, res) => {
  try {
    const item = Item.findById(req.params.id);
    
    if (!item) {
      return res.status(404).json({
        success: false,
        error: 'Item not found'
      });
    }
    
    res.status(200).json({
      success: true,
      data: item
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
};

// Create new item
const createItem = (req, res) => {
  try {
    const { name, description, price, quantity } = req.body;
    
    // Validation
    if (!name || !price) {
      return res.status(400).json({
        success: false,
        error: 'Name and price are required'
      });
    }
    
    const newItem = Item.create({
      name,
      description: description || '',
      price: parseFloat(price),
      quantity: parseInt(quantity) || 0
    });
    
    res.status(201).json({
      success: true,
      data: newItem
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
};

// Update item
const updateItem = (req, res) => {
  try {
    const { name, description, price, quantity } = req.body;
    
    const updatedItem = Item.update(req.params.id, {
      name,
      description,
      price: price ? parseFloat(price) : undefined,
      quantity: quantity ? parseInt(quantity) : undefined
    });
    
    if (!updatedItem) {
      return res.status(404).json({
        success: false,
        error: 'Item not found'
      });
    }
    
    res.status(200).json({
      success: true,
      data: updatedItem
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
};

// Delete item
const deleteItem = (req, res) => {
  try {
    const deleted = Item.delete(req.params.id);
    
    if (!deleted) {
      return res.status(404).json({
        success: false,
        error: 'Item not found'
      });
    }
    
    res.status(200).json({
      success: true,
      message: 'Item deleted successfully'
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
};

module.exports = {
  getAllItems,
  getItemById,
  createItem,
  updateItem,
  deleteItem
};
