## This script generates a file called tls.log. The difference from ssl.log is that this
## is much more focused on logging all kinds of protocol features. This can be interesting
## for academic purposes - or if one is just interested in more information about specific
## features used in local TLS traffic.
##@load base/protocols/ssl
##@load base/protocols/conn
module TLSLog;

export {
	## Log identifier for certificate log, as well as for the connection information log
	redef enum Log::ID += {
		TLS_LOG,
	};
	
	type TLSInfo: record {
		## Timestamp when the conenction began.
		pcap: string &log;
		ts: time &log;
		## Connection uid
		uid: string &log;
		## Connection 4-tuple
		id: conn_id &log;
		duration: interval &log;
		or_spl: vector of int &log &optional;

		## Numeric version of the server in the server hello
		server_version: count &log &optional;
		## Numeric version of the client in the client hello
		client_version: count &log &optional;
		## Cipher that was chosen for the connection
		cipher: count &log &optional;
		## Ciphers that were offered by the client for the connection
		client_ciphers: vector of count  &log &optional;
		## SNI that was sent by the client
		sni: vector of string &log &optional;
		## SSL Client extensions
		ssl_client_exts: vector of count &log &optional;
		## SSL server extensions
		ssl_server_exts: vector of count &log &optional;

		## Set to true if the ssl_established event was seen.
		ssl_established: bool &log &default=F;
		## TLS 1.3 supported versions
		client_supported_versions: vector of count &log &optional;
		## TLS 1.3 supported versions
		server_supported_version: count &log &optional;

		issuer: string &log &optional;
		subject: string &log &optional;
		cert_num: count &log &optional;
		not_before: time &log &optional;
		not_after: time &log &optional;
		san_num: count &log &optional;
		ext_num: count &log &optional;

		## Suggested ticket lifetime sent in the session ticket handshake
		## by the server.
		#ticket_lifetime_hint: count &log &optional;
		## Hashes of the full certificate chain sent by the server
		#server_certs: vector of string &log &optional;
		## Hashes of the full certificate chain sent by the server
		#client_certs: vector of string &log &optional;
		
		## The diffie helman parameter size, when using DH.
		#dh_param_size: count &log &optional;
		## supported elliptic curve point formats
		#point_formats: vector of count  &log &optional;
		## The curves supported by the client.
		#client_curves: vector of count  &log &optional;
		## The curve the server chose when using ECDH.
		#curve: count &log &optional;
		## Application layer protocol negotiation extension sent by the client.
		#orig_alpn: vector of string &log &optional;
		## Application layer protocol negotiation extension sent by the server.
		#resp_alpn: vector of string &log &optional;
		## Alert. If the connection was closed with an TLS alers before being
		## completely established, this field contains the alert level and description
		## numbers that were transfered
		#alert: vector of count  &log &optional;
		## TLS 1.3 Pre-shared key exchange modes
		#psk_key_exchange_modes: vector of count &log &optional;
		## Key share groups from client hello
		#client_key_share_groups: vector of count &log &optional;
		## Selected key share group from server hello
		#server_key_share_group: count &log &optional;
		## Client supported compression methods
		#client_comp_methods: vector of count &log &optional;
		## Server chosen compression method
		#comp_method: count;
		## Client supported signature algorithms
		#sigalgs: vector of count &log &optional;
		## Client supported hash algorithms
		#hashalgs: vector of count &log &optional;
	};

}

redef record connection += {
	tls_conns: TLSInfo &optional;
};

global pcap_name:string;
event zeek_init() &priority=5
	{
	Log::create_stream(TLSLog::TLS_LOG, [$columns=TLSInfo, $path="tls"]);
	local source=packet_source();
	pcap_name=source$path;
	}

event zeek_init()
    {
	Log::remove_filter(SSL::LOG, "default");
	Log::remove_filter(Conn::LOG, "default");
	Log::remove_filter(X509::LOG, "default");
	Log::remove_filter(PacketFilter::LOG, "default");
	Log::remove_filter(Files::LOG, "default");
	Log::remove_filter(OCSP::LOG, "default");
	local dir = @DIR;
    local filter: Log::Filter =
        [
        $name="sqlite",
        $path=dir+"/fivetuple",
        $config=table(["tablename"] = "fivetuple"),
        $writer=Log::WRITER_SQLITE
        ];
     Log::add_filter(TLSLog::TLS_LOG, filter);
    }

function set_session(c: connection)
	{
	if ( ! c?$tls_conns )
		{
		local t: TLSInfo;
		t$pcap=pcap_name;
		t$ts=network_time();
		t$uid=c$uid;
		t$id=c$id;
		t$ssl_client_exts=vector();
		t$ssl_server_exts=vector();
		t$or_spl = vector();
		c$tls_conns = t;
		}
	}

event tcp_packet(c: connection, is_orig: bool, flags: string, seq: count, ack: count, len: count, payload: string) {
    if(len != 0) {
		set_session(c);

        if( is_orig==T) {
            c$tls_conns$or_spl += len * (-1);
        } else {
            c$tls_conns$or_spl += len;
        }
    }
}

event ssl_client_hello(c: connection, version: count, record_version: count, possible_ts: time, client_random: string, session_id: string, ciphers: index_vec, comp_methods: index_vec)
	{
	set_session(c);
	c$tls_conns$client_ciphers = ciphers;
	c$tls_conns$client_version = version;
	}

event ssl_server_hello(c: connection, version: count, record_version: count, possible_ts: time, server_random: string, session_id: string, cipher: count, comp_method: count)
	{
	set_session(c);
	c$tls_conns$server_version = version;
	c$tls_conns$cipher = cipher;
	}

event ssl_extension(c: connection, is_orig: bool, code: count, val: string)
	{
	set_session(c);

	if ( is_orig )
		c$tls_conns$ssl_client_exts[|c$tls_conns$ssl_client_exts|] = code;
	else
		c$tls_conns$ssl_server_exts[|c$tls_conns$ssl_server_exts|] = code;
	}

event ssl_extension_server_name(c: connection, is_orig: bool, names: string_vec)
	{
	set_session(c);
	if ( !is_orig )
		return;

	c$tls_conns$sni = names;
	}

event ssl_established(c: connection)
	{
	set_session(c);

	c$tls_conns$ssl_established = T;
	}


event ssl_extension_supported_versions(c: connection, is_orig: bool, versions: index_vec)
	{
	set_session(c);
	if ( is_orig )
		c$tls_conns$client_supported_versions = versions;
	else
		c$tls_conns$server_supported_version = versions[0];
	}

function log_cert_chain(c: connection, chain: vector of Files::Info)
	{
		local ser_cert= chain[0];
		if ( !ser_cert?$x509 || !ser_cert$x509?$handle )
		{
			return;
		}
		if (c$ssl?$issuer){
			c$tls_conns$issuer = c$ssl$issuer;
		}
		if (c$ssl?$subject){
			c$tls_conns$subject = c$ssl$subject;
		}
		c$tls_conns$not_before = ser_cert$x509$certificate$not_valid_before;
		c$tls_conns$not_after =ser_cert$x509$certificate$not_valid_after;

		c$tls_conns$cert_num= |chain|;
		
		if(ser_cert$x509?$extensions) {
			c$tls_conns$ext_num = |ser_cert$x509$extensions|;
		}
		if(ser_cert$x509?$san) {
			local san_num = 0;
			if(ser_cert$x509$san?$dns){
				san_num+=|ser_cert$x509$san$dns|;
			}
			if(ser_cert$x509$san?$uri){
				san_num+=|ser_cert$x509$san$uri|;
			}
			if(ser_cert$x509$san?$email){
				san_num+=|ser_cert$x509$san$email|;
			}
			if(ser_cert$x509$san?$ip){
				san_num+=|ser_cert$x509$san$ip|;
			}
			c$tls_conns$san_num = san_num;
		}
		
	}

event connection_state_remove(c: connection)
	{
	if ( ! c?$ssl)
		return;
	if ( ! c?$tls_conns )
		return;

	if ( c$ssl?$cert_chain )
		log_cert_chain(c, c$ssl$cert_chain);
	c$tls_conns$duration = c$duration;
	Log::write(TLSLog::TLS_LOG, c$tls_conns);
	}


