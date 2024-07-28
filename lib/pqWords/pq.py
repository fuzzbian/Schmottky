import sys
import sympy as sp
import numpy as np

# TODO: tidy me up!

def get_trace_polynomial(p,q,t_a,t_B,t_aB):
# of the word p/q in variables t_a, t_B, t_aB
    trace_poly=[]
    
    if (p,q) == (0,1):
        return t_a
    if (p,q) == (1,0):
        return t_B

    (p1,q1) = (0,1)
    (p2,q2) = (1,0)
    (p3,q3) = (1,1)

    t_u = t_a
    t_v = t_B
    t_uv = t_aB

    while not np.isclose(p3/q3,p/q):
        
        if p/q < p3/q3:
            # go left
            (p2,q2) = (p3,q3)
            (p3,q3) = (p1+p3,q1+q3)
            
            temp = t_uv
            
            t_uv = t_u * t_uv - t_v
            t_v = temp

        else:
            # go right
            (p1,q1) = (p3,q3)
            (p3,q3) = (p2+p3,q2+q3)
            
            temp = t_uv
            
            t_uv = t_v * t_uv - t_u
            t_u = temp

    return t_uv


def farey(q_max):
    # inital sequece
    sequence = [(0,1),(1,1)]
    # farey addition
    f_add = lambda x, y: (x[0]+y[0],x[1]+y[1])
    N = 1
    while N < q_max:
        N += 1
        # a beautiful home for our new fractions
        new_subsequence = []
        # generate new fractions
        for i in range(len(sequence)-1):
            new = f_add(sequence[i], sequence[i+1])
            if not new[1] > N:
                new_subsequence.append(new)
        # add them to the sequence and bring order to the mess we made by adding new shit
        sequence += new_subsequence
        sequence.sort(key=lambda element: element[0]/element[1])

    return sequence


def get_taB_with_markov(t_a,t_B):
# for given traces t_a and t_B returns t_aB so that it fulfills the Markov identity
# i.e. t_abAB = -2
# if used with numbers, they must be complex
    
    return t_a*t_B/2 + sp.sqrt((t_a*t_B/2)**2-t_a**2-t_B**2)


def get_mu(pq):

    # variables
    t_B = complex(2)
    t_a, t_aB = sp.symbols('t_a t_aB')

    # get t_aB as a polynamial in t_a
    t_aB = get_taB_with_markov(t_a,t_B)
    solutions = []
    mu0 = 0 + 2j
    
    for p, q in farey(pq[1]):
        # get trace polynomial 
        t_pq = get_trace_polynomial(p,q,t_a,t_B,t_aB)
        t_pq = sp.Poly(t_pq,t_a)
        #print("TRACE POLYNOMIAL",t_pq)
        #print("DEGREE", t_pq.degree_list())

        t_pq = t_pq - 2
        #t_pq2 = t_pq + 2
        solutionspq = t_pq.nroots(maxsteps=100)
        nearest = solutionspq[0]
        for x in solutionspq:
            if np.abs(x-mu0) <= np.abs(nearest-mu0):
                nearest = x
        #if (p,q) == (1,10):
        #    print('MUUUUUUUUUUUUUUUUU 1/10',nearest)
        solutions.append(nearest)
        mu0 = nearest
        if p == pq[0] and q == pq[1]:
            return mu0

p = int(sys.argv[1])
q = int(sys.argv[2])

mu = complex(get_mu((p,q)))
print(f"{mu.real};{mu.imag}",end='')