import ipaddress
import time

lookup= b'' +\
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00' +\
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00' +\
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00' +\
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00' +\
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00' +\
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00' +\
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00' +\
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00' +\
b'\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01' +\
b'\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01' +\
b'\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01' +\
b'\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01' +\
b'\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02' +\
b'\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02' +\
b'\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03' +\
b'\x04\x04\x04\x04\x04\x04\x04\x04\x05\x05\x05\x05\x06\x06\x07\x08'

def bitmask_num(x):
    m = lookup[x>>24]
    if m != 8:
        return m

    v = (x << 8) & 0xffffffff 
    i = 0
    while i < 3:
        lk= lookup[v>>24]
        m += lk
        if lk != 8:
            return m
        v = (v<<8) & 0xffffffff
        i +=  1
    return m
    
# this on only for 2 set of ip, as your example
def calc_subnet(ip1, ip2):
    ipi1 = int(ipaddress.IPv4Address(ip1))
    ipi2 = int(ipaddress.IPv4Address(ip2))

    bitmask=~( ipi1 ^ ipi2)
    return ipaddress.IPv4Network(str(ipaddress.IPv4Address(ipi1)) + '/' + str(bitmask_num(bitmask)), strict=False)

def calc_subnets(*ips):
    ip = int(ipaddress.IPv4Address(ips[0]))
    bitmask = 0xffffffff

    i=1
    if len(ips) > 1:
        while i < len(ips):
            bitmask &= 0xffffffff & ~(ip ^ (int(ipaddress.IPv4Address(ips[i])))) 
            ip = ip & bitmask
            i += 1
    bitnum = bitmask_num(bitmask)
    mask = (0xffffffff << 32-bitnum) & 0xffffffff
    ip &= mask
#     return ipaddress.IPv4Network(str(ipaddress.IPv4Address(ip)) + '/' + str(bitnum))

    '''
    if return is only string (instead of ipaddress.IP4vNetwork, this is way much faster
    Running 1000,000 loops
    Using  original finished in :  30.593202829360962
    Using  min max  finished in :  14.5304274559021
    Using  bitwise finished in  :  13.740501642227173
    '''
    return str(ipaddress.IPv4Address(ip)) + '/' + str(bitnum)

def calc_subnets_minmax(*ips):
    min = 0xffffffff
    max = 0
    i = 0
    while i < len(ips):
        v = int(ipaddress.IPv4Address(ips[i]))
        i += 1
        if v < min:
            min = v
            continue
        if v > max:
            max = v
    if max <= min:
        return calc_subnets(min)
    return calc_subnets(min, max)
'''
THIS IS YOUR ORIGINAL CODE
'''
#accepts 2 IP strings
def calc_subnet_original(ip1, ip2):

    #define IP Address objects
    ip1_obj=ipaddress.IPv4Address(ip1)
    ip2_obj=ipaddress.IPv4Address(ip2)
    
    if ip1_obj<=ip2_obj:
        min_ip=ip1_obj
        max_ip=ip2_obj
    else:
        min_ip=ip2_obj
        max_ip=ip1_obj
        
    distance = int(max_ip)-int(min_ip)
    ip_range=0 

    # increase power of 2 until we find subnet distance
    while 2**ip_range < distance:
        ip_range += 1
          
    net = ipaddress.IPv4Network(str(min_ip) + '/' +str(32-ip_range), strict=False)
    if max_ip not in net: 

    # if the distance implies one size network, but IPs span 2
        ip_range+=1
        net = ipaddress.IPv4Network(str(min_ip) + '/' +str(32-ip_range), strict=False)
        
    return net

''' 
MORE THAN 1 IPs
'''
def calc_subnets_original(*ips):
    min = ipaddress.IPv4Address(ips[0])
    max = min 

    i = 0
    while i<len(ips):
        v = ipaddress.IPv4Address(ips[i])
        i += 1
        if v > max:
            max = v
            continue
        if v < min:
            min = v

    distance = int(max) - int(min)
    ip_range = 0

    while 2**ip_range < distance:
        ip_range += 1

    net = ipaddress.IPv4Network(str(min) + '/' +str(32-ip_range), strict=False)
    if max not in net: 
    # if the distance implies one size network, but IPs span 2
        ip_range+=1
        net = ipaddress.IPv4Network(str(min) + '/' +str(32-ip_range), strict=False)
        
    return net

# print('--- original data test ---')
# print()
# print('DATA:','192.168.0.1','192.168.0.128')
# print('original:',calc_subnet_original('192.168.0.1','192.168.0.128'))
# print('bitwise :', calc_subnet('192.168.0.1','192.168.0.128'))
# print()
# print()
# print('DATA:','192.168.0.1','172.0.0.1')
# print('original:',calc_subnet_original('192.168.0.1','172.0.0.1'))
# print('bitwise :', calc_subnet('192.168.0.1','172.0.0.1'))
# print()
# print()
#
# print('--- more than 2 data for test ---')
# print()
# print('DATA:','192.168.1.192','192.168.1.128','192.168.1.133','192.168.1.100')
# print('original:',calc_subnets_original('192.168.1.192','192.168.1.128','192.168.1.133','192.168.1.100'))
# print('min max :',calc_subnets_minmax('192.168.1.192','192.168.1.128','192.168.1.133','192.168.1.100'))
# print('bitsiwe :',calc_subnets('192.168.1.192','192.168.1.128','192.168.1.133','192.168.1.100'))
# print()
# print('DATA:','192.168.1.192','192.168.1.128','192.168.1.133','172.1.1.1')
# print('original:',calc_subnets_original('192.168.1.192','192.168.1.128','192.168.1.133','172.1.1.1'))
# print('min max :',calc_subnets_minmax('192.168.1.192','192.168.1.128','192.168.1.133','172.1.1.1'))
# print('bitwise :',calc_subnets('192.168.1.192','192.168.1.128','192.168.1.133','172.1.1.1'))
# print()
# print()



print('Python Running 100,000 loops')
loop = 0
maxloop = 100000
t1 = time.time()
testcase=['192.168.1.192','192.168.1.128','192.168.1.133','192.168.1.100']
print('DATA:', testcase)
# while loop < maxloop:
#     calc_subnets_original(*testcase)
#     loop += 1
#
# t2 = time.time()
# print("Using  original finished in : ", t2-t1)

loop = 0
t1 = time.time()
while loop < maxloop:
    calc_subnets_minmax(*testcase)
    loop += 1

t2 = time.time()
print("Using  min max  finished in : ", t2-t1)
print("     Result                 : ", calc_subnets_minmax(*testcase))


loop = 0
t1 = time.time()
while loop < maxloop:
    calc_subnets(*testcase)
    loop += 1

t2 = time.time()

loop = 0
t1 = time.time()
while loop < maxloop:
    calc_subnets(*testcase)
    loop += 1

t2 = time.time()
print("Using  bitwise finished in  : ", t2-t1)
print("     Result                 : ", calc_subnets(*testcase))

print()
loop=0
t1 = time.time()
testcase=['192.168.244.255','172.31.255.255','172.30.1.1','192.168.1.1','192.168.1.192','192.168.1.128','192.168.1.133','192.168.1.100','172.16.1.1','128.1.1.90','10.1.1.1', '10.2.2.2']
print('DATA:', testcase)
# while loop < maxloop:
#     calc_subnets_original(*testcase)
#     loop += 1
# # print(calc_subnets_original(*testcase))
# t2 = time.time()
# print("Using  original finished in : ", t2-t1)

loop = 0
t1 = time.time()
while loop < maxloop:
    calc_subnets_minmax(*testcase)
    loop += 1
# print(calc_subnets_minmax(*testcase))
t2 = time.time()
print("Using  min max  finished in : ", t2-t1)
print("     Result                 : ", calc_subnets_minmax(*testcase))


loop = 0
t1 = time.time()
while loop < maxloop:
    calc_subnets(*testcase)
    loop += 1
# print(calc_subnets(*testcase))
t2 = time.time()

print("Using  bitwise finished in  : ", t2-t1)
print("     Result                 : ", calc_subnets(*testcase))
