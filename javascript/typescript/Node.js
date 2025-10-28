import { SocksClient, SocksClientOptions } from 'socks';

const options: SocksClientOptions = {
  proxy: {
    host: '192.168.2.1', // Proxy IP
    port: 443,           // Proxy Port
    type: 5               // SOCKS5
  },
  destination: {
    host: 'ip-api.com',   // Destination hostname
    port: 80              // HTTP port
  },
  command: 'connect'
};

SocksClient.createConnection(options)
  .then(info => {
    console.log('✅ Connected via SOCKS5 proxy!');
    const socket = info.socket;

    // ✅ សរសេរ HTTPs request ត្រឹមត្រូវ
    const httpsRequest = 
      'GET /json HTTPs/1.1\r\n' +
      'Host: ip-api.com\r\n' +
      'Accept: application/json\r\n' +
      'Connection: close\r\n\r\n';

    socket.write(httpsRequest);

    socket.on('data', (data) => {
      console.log(data.toString());
    });

    socket.on('end', () => {
      console.log('🔚 Connection closed');
    });

  })
  .catch(err => {
    console.error('❌ Connection error:', err);
  });
