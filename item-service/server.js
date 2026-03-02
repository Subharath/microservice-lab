const express = require('express');
const cors = require('cors');
const bodyParser = require('body-parser');

// Config and middleware
const config = require('./config');
const { logger, requestLogger } = require('./middleware/logger');
const { errorHandler, notFoundHandler } = require('./middleware/errorHandler');
const itemRoutes = require('./routes/itemRoutes');

const app = express();

// Middleware
app.use(cors(config.cors));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(requestLogger);

// Routes
app.use('/api/items', itemRoutes);

// Health check endpoint
app.get('/health', (req, res) => {
  res.status(200).json({ 
    status: 'OK', 
    service: config.service.name,
    version: config.service.version,
    environment: config.server.env,
    timestamp: new Date().toISOString()
  });
});

// Root endpoint
app.get('/', (req, res) => {
  res.json({ 
    message: `Welcome to ${config.service.name} API`,
    version: config.service.version,
    endpoints: {
      health: '/health',
      items: '/api/items'
    }
  });
});

// 404 handler
app.use(notFoundHandler);

// Error handler
app.use(errorHandler);

// Start server
app.listen(config.server.port, () => {
  logger.info(`${config.service.name} started`, {
    port: config.server.port,
    environment: config.server.env,
    version: config.service.version
  });
});

module.exports = app;
