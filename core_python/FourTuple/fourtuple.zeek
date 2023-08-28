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
		pcap: string &log;
		## Timestamp when the conenction began.
		ts: time &log;
		## Connection uid
		uid: string &log;
		## Connection 4-tuple
		id: conn_id &log;
		duration: interval &log;

		up_bytes: int &log; # new
		down_bytes: int &log;  # new
		up_pkts: count &log; # new
		down_pkts: count &log; # new

		## Numeric version of the server in the server hello
		server_version: count &log &optional;
		## Numeric version of the client in the client hello
		client_version: count &log &optional;
		## SNI that was sent by the client
		sni: vector of string &log &optional;

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
		valid_time: vector of interval &log &optional; # new
		cert_serial: vector of string &log &optional; # new
		cert_len: vector of int &log &optional; # new
		san_num: count &log &optional;
		ext_num: count &log &optional;

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
	Log::remove_filter(HTTP::LOG, "default");
	Log::remove_filter(DNS::LOG, "default");
	Log::remove_filter(PacketFilter::LOG, "default");
	local dir = @DIR;
    local filter: Log::Filter =
        [
        $name="sqlite",
        $path=dir+"/fourtuple",
        $config=table(["tablename"] = "fourtuple"),
        $writer=Log::WRITER_SQLITE
        ];
    Log::add_filter(TLSLog::TLS_LOG, filter);
    }

function set_session(c: connection)
	{
	if ( ! c?$tls_conns )
		{
		local t: TLSInfo;
		t$ts=network_time();
		t$pcap=pcap_name;
		t$uid=c$uid;
		t$id=c$id;

		t$up_bytes=0;
		t$down_bytes=0;
		t$up_pkts=0;
		t$down_pkts=0;
		t$cert_serial=vector();
		t$valid_time=vector();
		t$cert_len=vector();

		c$tls_conns = t;
		}
	}


event ssl_client_hello(c: connection, version: count, record_version: count, possible_ts: time, client_random: string, session_id: string, ciphers: index_vec, comp_methods: index_vec)
	{
	set_session(c);
	c$tls_conns$client_version = version;
	}

event ssl_server_hello(c: connection, version: count, record_version: count, possible_ts: time, server_random: string, session_id: string, cipher: count, comp_method: count)
	{
	set_session(c);
	c$tls_conns$server_version = version;
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
		for (i in chain)
		{
			local cert = chain[i];
			c$tls_conns$cert_serial += cert$x509$certificate$serial;
			c$tls_conns$valid_time += (cert$x509$certificate$not_valid_after-cert$x509$certificate$not_valid_before);
			c$tls_conns$cert_len += cert$seen_bytes;
		}
		
		local ser_cert= chain[0];
		if ( !ser_cert?$x509 || !ser_cert$x509?$handle )
		{
			return;
		}
		if (c$ssl?$issuer)
		{
			c$tls_conns$issuer = c$ssl$issuer;
		}
		if (c$ssl?$subject) 
		{		
			c$tls_conns$subject = c$ssl$subject;
		}
		c$tls_conns$not_before = ser_cert$x509$certificate$not_valid_before;
		c$tls_conns$not_after =ser_cert$x509$certificate$not_valid_after;
		# print ser_cert$x509$certificate$not_valid_after;


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
	# print c$id$orig_h;
	# print c$orig$num_pkts,c$resp$num_pkts;
	if ( ! c?$tls_conns )
		return;
	c$tls_conns$up_bytes = c$orig$size;
	c$tls_conns$up_pkts = c$orig$num_pkts;

	c$tls_conns$down_bytes = c$resp$size;
	c$tls_conns$down_pkts = c$resp$num_pkts;

	if ( ! c?$ssl)
		return;
	if ( c$ssl?$cert_chain )
		log_cert_chain(c, c$ssl$cert_chain);
	c$tls_conns$duration = c$duration;
	Log::write(TLSLog::TLS_LOG, c$tls_conns);
	}


